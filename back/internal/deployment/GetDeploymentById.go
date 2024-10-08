package deployment

import (
	"context"
	"fmt"

	"github.com/nallanos/fire/internal/db"
)

func (c *ContainerService) GetDeploymentById(ctx context.Context, deployment_id string) (db.Deployment, error) {
	deployment, err := c.db.GetDeploymentById(ctx, deployment_id)
	if err != nil {
		return deployment, fmt.Errorf("error getting deployment: %w", err)
	}
	return deployment, nil
}
