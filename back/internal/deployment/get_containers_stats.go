package deployment

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
)

type StatsData struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
}

func (c *ContainerService) GetContainerStats(ctx context.Context, app_id string) (types.StatsJSON, error) {

	container, err := c.ListApplicationContainers(ctx, app_id)
	if err != nil {
		return types.StatsJSON{}, fmt.Errorf("error while getting listing application containers %w", err)
	}

	stats, err := c.docker.ContainerStats(ctx, container[0].ID, true)
	if err != nil {
		return types.StatsJSON{}, fmt.Errorf("error while getting container stats: %w", err)
	}
	defer stats.Body.Close()

	var containerStats types.StatsJSON
	if err := json.NewDecoder(stats.Body).Decode(&containerStats); err != nil && err != io.EOF {
		return types.StatsJSON{}, fmt.Errorf("error while decoding container stats: %w", err)
	}

	return containerStats, nil
}
