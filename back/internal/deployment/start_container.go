package deployment

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/nallanos/fire/internal/db"
)

func (s *ContainerService) StartContainer(app *db.Application, deployment db.Deployment) error {
	containers, err := s.ListApplicationContainers(context.Background(), app.ID)
	if err != nil {
		return fmt.Errorf("error listing application containers: %w", err)
	}
	err = s.docker.ContainerStart(context.Background(), containers[0].ID, container.StartOptions{})
	if err != nil {
		s.db.UpdateDeployment(context.Background(), db.UpdateDeploymentParams{
			FinishedAt: sql.NullTime{Time: time.Now(), Valid: true},
			ID:         deployment.ID,
			Status:     StatusFailed,
		})
		return fmt.Errorf("error starting container: %w", err)
	}
	err = s.db.UpdateDeployment(context.Background(), db.UpdateDeploymentParams{
		FinishedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:         deployment.ID,
		Status:     StatusActive,
	})
	if err != nil {
		return fmt.Errorf("error updating deployment: %w", err)
	}
	err = s.db.UpdateApplication(context.Background(), db.UpdateApplicationParams{
		ID:     app.ID,
		Image:  app.Image,
		Port:   app.Port,
		Name:   app.Name,
		Status: StatusActive,
	})
	if err != nil {
		return fmt.Errorf("error updating app: %w", err)
	}
	return nil
}
