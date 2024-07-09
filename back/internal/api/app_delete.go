package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/go-chi/chi/v5"
)

func (a *API) deleteApp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println(id)
	app, err := a.apps.GetApplication(context.Background(), id)
	if err != nil {
		slog.Error("error getting1 application: %w", err)
	}
	err = a.apps.DeleteApplication(r.Context(), id)
	if err != nil {
		http.Error(w, "error while deleting application", http.StatusNotFound)
		InternalServerError(w, err)
		return
	}
	containers, err := a.deployments.ListApplicationContainers(context.Background(), app.ID)
	if err != nil {
		slog.Error("error listing containers: %w", err)
	}

	for _, c := range containers {
		if err := a.docker.ContainerStop(context.Background(), c.ID, container.StopOptions{}); err != nil {
			slog.Error("error stopping container: %w", err)
		}

		if err := a.docker.ContainerRemove(context.Background(), c.ID, container.RemoveOptions{}); err != nil {
			slog.Error("error removing container: %w", err)
		}
	}
}
