package deployment

import (
	"context"
	"fmt"

	"github.com/nallanos/fire/internal/db"
)

func (c *ContainerService) ListApplicationDeployments(ctx context.Context, id string) ([]db.Deployment, error) {
	deployments, err := c.db.ListApplicationDeployments(ctx, id)

	if err != nil {
		return deployments, fmt.Errorf("error listing application deployments: %w", err)
	}

	return deployments, nil
}
