package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *API) stopContainer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	app, err := a.apps.GetApplication(context.Background(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		InternalServerError(w, err)
		return
	}

	deployment, err := a.deployments.GetLatestDeployment(context.Background(), id)

	if err != nil {
		slog.Error("error getting latest deployment", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		InternalServerError(w, err)
		return
	}
	err = a.deployments.StopContainer(app, deployment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		InternalServerError(w, err)
		return
	}
}
