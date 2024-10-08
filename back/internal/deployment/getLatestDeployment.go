package deployment

import (
	"context"
	"fmt"

	"github.com/nallanos/fire/internal/db"
)

func (s *ContainerService)GetLatestDeployment(ctx context.Context, app_id string) (db.Deployment, error) {
	deployment, err := s.db.GetLatestDeployment(context.Background(), app_id)
	if err != nil {
		return deployment, fmt.Errorf("error getting deployment: %w", err)
	}
	return deployment, nil
}
