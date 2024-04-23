package deployment

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/docker/docker/api/types/container"
	"github.com/nallanos/fire/internal/db"
)

func (s *ContainerService) StopContainer(app *db.Application) error {
	containers, err := s.ListApplicationContainers(context.Background(), app.ID)
	if err != nil {
		slog.Error("error stopping the Container", err)
		return fmt.Errorf("error listing the Container %w", err)
	}
	err = s.docker.ContainerStop(context.Background(), containers[0].ID, container.StopOptions{})
	if err != nil {
		slog.Error("error stopping the Container", err)
		return fmt.Errorf("error stopping the Container %w", err)
	}
	return nil
}
