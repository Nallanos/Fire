package deployment

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/nallanos/fire/internal/db"
	"github.com/nrednav/cuid2"
)

type Service struct {
	docker *client.Client
	db     *db.Queries
}

func NewService(docker *client.Client, db *db.Queries) *Service {
	return &Service{
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

func (s *Service) DeployApplication(app *db.Application) error {
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

	containers, err := s.docker.ContainerList(context.Background(), types.ContainerListOptions{
		Filters: filters.NewArgs(filters.Arg("label", "app_id="+app.ID)),
	})
	if err != nil {
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

	reader, err := s.docker.ImagePull(context.Background(), app.Image, types.ImagePullOptions{})
	if err != nil {
		slog.Error("Error pulling image", "err", err)
		return fmt.Errorf("error pulling image: %w", err)
	}

	slog.Info("Copying image pull output")
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		return fmt.Errorf("error downloading image : %w", err)
	}

	container, err := s.docker.ContainerCreate(context.Background(),
		&container.Config{
			Labels: map[string]string{
				"app_id": app.ID,
			},
			Image: app.Image,
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
		}, nil, nil, app.ID)
	if err != nil {
		return fmt.Errorf("error creating container: %w", err)
	}

	slog.Info("Starting container", "id", container.ID)
	err = s.docker.ContainerStart(context.Background(), container.ID, types.ContainerStartOptions{})
	if err != nil {
		slog.Error("Error starting container", "err", err)
		return fmt.Errorf("error starting container: %w", err)
	}

	if err := s.db.UpdateDeployment(context.Background(), db.UpdateDeploymentParams{
		ID:     deployment.ID,
		Status: StatusCompleted,
	}); err != nil {
		return fmt.Errorf("error updating deployment in db: %w", err)
	}

	return nil
}
