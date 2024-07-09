package deployment

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/nallanos/fire/internal/db"
	"github.com/nrednav/cuid2"
)

type ContainerService struct {
	docker *client.Client
	db     *db.Queries
}

func NewService(docker *client.Client, db *db.Queries) *ContainerService {
	return &ContainerService{
		docker: docker,
		db:     db,
	}
}

type Status = string

const (
	StatusPending  Status = "pending"
	StatusFailed   Status = "failed"
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

func (s *ContainerService) Deploy(app *db.Application) error {
	deployment, err := s.db.CreateDeployment(context.Background(), db.CreateDeploymentParams{
		Status: StatusPending,
		AppID:  app.ID,
		ID:     cuid2.Generate(),
	})
	if err != nil {
		return fmt.Errorf("error creating deployment in db: %w", err)
	}

	defer func() {
		if err != nil {
			s.db.UpdateDeployment(context.Background(), db.UpdateDeploymentParams{
				Status: StatusFailed,
			})
			s.db.UpdateApplication(context.Background(), db.UpdateApplicationParams{
				Status: StatusFailed,
			})
		}
	}()

	containers, err := s.ListApplicationContainers(context.Background(), app.ID)
	if err != nil {
		return fmt.Errorf("error listing containers: %w", err)
	}

	for _, c := range containers {
		if err := s.docker.ContainerStop(context.Background(), c.ID, container.StopOptions{}); err != nil {
			return fmt.Errorf("error stopping container: %w", err)
		}

		if err := s.docker.ContainerRemove(context.Background(), c.ID, container.RemoveOptions{}); err != nil {
			return fmt.Errorf("error removing container: %w", err)
		}
	}

	reader, err := s.docker.ImagePull(context.Background(), "docker.io/library/nginx:latest", types.ImagePullOptions{})
	if err != nil {
		slog.Error("error pulling image")
		return fmt.Errorf("error pulling image: %w", err)
	}

	if _, err := io.Copy(os.Stdout, reader); err != nil {
		return fmt.Errorf("error downloading image : %w", err)
	}

	container, err := s.docker.ContainerCreate(context.Background(),
		&container.Config{
			Labels: map[string]string{
				"app_id": app.ID,
			},
			Image: "docker.io/library/nginx:latest",
		}, nil, nil, nil, deployment.ID)
	if err != nil {
		return fmt.Errorf("error creating container: %w", err)
	}
	slog.Info("container created", "id", container.ID)

	err = s.StartContainer(app, deployment)

	if err != nil {
		return fmt.Errorf("error starting container: %w", err)
	}

	err = s.db.UpdateApplication(context.Background(), db.UpdateApplicationParams{
		ID:     app.ID,
		Image:  app.Image,
		Port:   app.Port,
		Name:   app.Name,
		Status: StatusActive,
	})

	if err != nil {
		return fmt.Errorf("error updating deployment in db: %w", err)
	}

	return nil
}
