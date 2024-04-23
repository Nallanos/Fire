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
	"github.com/docker/go-connections/nat"
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
	StatusStarded   Status = "started"
	StatusFailed    Status = "failed"
	StatusCompleted Status = "completed"
)

func (s *ContainerService) Deploy(app *db.Application) error {
	deployment, err := s.db.CreateDeployment(context.Background(), db.CreateDeploymentParams{
		AppID:  app.ID,
		ID:     cuid2.Generate(),
		Status: StatusStarded,
	})
	if err != nil {
		return fmt.Errorf("error creating deployment in db: %w", err)
	}

	defer func() {
		if err != nil {
			s.db.UpdateDeployment(context.Background(), db.UpdateDeploymentParams{
				ID:     deployment.ID,
				Status: StatusFailed,
			})
		}
	}()

	containers, err := s.ListApplicationContainers(context.Background(), app.ID)
	if err != nil {
		slog.Error("error listing container")
		return fmt.Errorf("error listing containers: %w", err)
	}

	for _, c := range containers {
		if err := s.docker.ContainerStop(context.Background(), c.ID, container.StopOptions{}); err != nil {
			return fmt.Errorf("error stopping container: %w", err)
		}

		if err := s.docker.ContainerRemove(context.Background(), c.ID, types.ContainerRemoveOptions{}); err != nil {
			return fmt.Errorf("error removing container: %w", err)
		}
	}

	reader, err := s.docker.ImagePull(context.Background(), "docker.io/library/docker", types.ImagePullOptions{})
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
		},
		&container.HostConfig{
			PortBindings: nat.PortMap{
				nat.Port("8081/tcp"): []nat.PortBinding{
					{
						HostPort: "3000",
						HostIP:   "0.0.0.0",
					},
				},
			},
		}, nil, nil, deployment.ID)
	if err != nil {
		slog.Error("error creating container", err)
		return fmt.Errorf("error creating container: %w", err)
	}
	slog.Info("Creating container", "id", container.ID)

	if err := s.db.UpdateDeployment(context.Background(), db.UpdateDeploymentParams{
		ID:     deployment.ID,
		Status: StatusCompleted,
	}); err != nil {
		return fmt.Errorf("error updating deployment in db: %w", err)
	}

	return nil
}
