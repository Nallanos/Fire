package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *API) startContainer(w http.ResponseWriter, r *http.Request) {
	app_id := chi.URLParam(r, "id")
	app, err := a.apps.GetApplication(context.Background(), app_id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		InternalServerError(w, err)
		return
	}
	deployment, err := a.deployments.GetLatestDeployment(context.Background(), app_id)

	if err != nil {
		slog.Error(fmt.Sprintf("error getting latest deployment %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		InternalServerError(w, err)
		return
	}

	err = a.deployments.StartContainer(app, deployment)
	if err != nil {
		slog.Error(fmt.Sprintf("error starting container %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		InternalServerError(w, err)
		return
	}
}
