package deployment

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/nallanos/fire/internal/db"
)

func (s *ContainerService) StartContainer(id string) error {
	containers, err := s.ListApplicationContainers(context.Background(), id)
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
	app, err := s.db.GetApplication(context.Background(), id)
	if err != nil {
		return fmt.Errorf("error getting app: %w", err)
	}
	err = s.db.UpdateApplication(context.Background(), db.UpdateApplicationParams{
		ID:     app.ID,
		Image:  app.Image,
		Port:   app.Port,
		Name:   app.Name,
		Status: StatusCompleted,
	})
	if err != nil {
		return fmt.Errorf("error updating app: %w", err)
	}
	return nil
}
