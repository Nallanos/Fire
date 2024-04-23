package deployment

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
)

func (c *ContainerService) ListApplicationContainers(ctx context.Context, appId string) ([]types.Container, error) {
	containers, err := c.docker.ContainerList(ctx, container.ListOptions{
		Filters: filters.NewArgs(filters.Arg("label", "app_id="+appId)),
		All: true,
	})
	if err != nil {
		return nil, err
	}

	return containers, nil
}
