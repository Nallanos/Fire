package deployment

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/nallanos/fire/internal/db"
)

func (s *ContainerService) StopContainer(app *db.Application) error {
	containers, err := s.ListApplicationContainers(context.Background(), app.ID)
	if err != nil {
		return fmt.Errorf("error listing the Container %w", err)
	}
	err = s.docker.ContainerStop(context.Background(), containers[0].ID, container.StopOptions{})
	if err != nil {
		s.db.UpdateDeployment(context.Background(), db.UpdateDeploymentParams{
			Status: StatusFailed,
		})
		return fmt.Errorf("error stopping the Container %w", err)
	}
	err = s.db.UpdateDeployment(context.Background(), db.UpdateDeploymentParams{
		Status: StatusInactive,
	})
	if err != nil {
		return fmt.Errorf("error updating deployment in db: %w", err)
	}
	err = s.db.UpdateApplication(context.Background(), db.UpdateApplicationParams{
		ID:     app.ID,
		Image:  app.Image,
		Port:   app.Port,
		Name:   app.Name,
		Status: "inactive",
	})

	if err != nil {
		return fmt.Errorf("error updating application status in db: %w", err)
	}
	return nil
}
