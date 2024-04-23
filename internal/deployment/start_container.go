package deployment

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/docker/docker/api/types/container"
	"github.com/nallanos/fire/internal/db"
)

func (s *ContainerService) StartContainer(app *db.Application) error {
	containers, err := s.ListApplicationContainers(context.Background(), app.ID)
	if err != nil {
		return fmt.Errorf("error listing application containers: %w", err)
	}
	err = s.docker.ContainerStart(context.Background(), containers[0].ID, container.StartOptions{})
	if err != nil {
		slog.Error("Error starting container", "err", err)
		return fmt.Errorf("error starting container: %w", err)
	}
	return nil
}
