package deployment

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types/container"
)

func (c *ContainerService) GetContainerLogs(ctx context.Context, app_id string) (string, error) {
	containers, err := c.ListApplicationContainers(ctx, app_id)
	if err != nil {
		return "", fmt.Errorf("error while listing applications containers: %w", err)
	}
	logsReader, err := c.docker.ContainerLogs(ctx, containers[0].ID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		return "", fmt.Errorf("error while getting logs containers: %w", err)
	}
	defer logsReader.Close()

	var containerLogs string
	buf := new(strings.Builder)
	if _, err := io.Copy(buf, logsReader); err != nil {
		return "", fmt.Errorf("error while reading container logs: %w", err)
	}
	containerLogs = buf.String()
	return containerLogs, nil
}
