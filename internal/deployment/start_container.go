package deployment

import (
	"context"
	"fmt"

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
		s.db.UpdateDeployment(context.Background(), db.UpdateDeploymentParams{
			Status: StatusFailed,
		})
		return fmt.Errorf("error starting container: %w", err)
	}
	err = s.db.UpdateDeployment(context.Background(), db.UpdateDeploymentParams{
		Status: StatusCompleted,
	})
	if err != nil {
		return fmt.Errorf("error updating deployment: %w", err)
	}
	return nil
}
