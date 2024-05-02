package deployment

import (
	"context"
	"fmt"

	"github.com/nallanos/fire/internal/db"
)

func (s *ContainerService) GetLatestDeployment(ctx context.Context, id string) (db.Deployment, error) {
	deployment, err := s.db.GetLatestDeployment(context.Background(), id)
	if err != nil {
		return deployment, fmt.Errorf("error getting deployment: %w", err)
	}
	return deployment, nil
}
