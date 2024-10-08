package deployment

import (
	"context"
	"fmt"

	"github.com/nallanos/fire/internal/db"
)

/*
Return the list of deployment belonging to an applications
*/
func (c *ContainerService) ListApplicationDeployments(ctx context.Context, app_id string) ([]db.Deployment, error) {
	deployments, err := c.db.ListApplicationDeployments(ctx, app_id)

	if err != nil {
		return deployments, fmt.Errorf("error listing application deployments: %w", err)
	}

	return deployments, nil
}
