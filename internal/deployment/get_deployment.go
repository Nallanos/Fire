package deployment

import (
	"context"
	"fmt"

	"github.com/nallanos/fire/internal/db"
)

func (s *ContainerService) GetDeployment(id string, ctx context.Context) (db.Deployment, error) {
	deployment, err := s.db.GetDeployment(context.Background(), id)

	if err != nil {
		return deployment, fmt.Errorf("error getting deployments: %w", err)
	}

	return deployment, nil
}
