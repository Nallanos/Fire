package api

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nallanos/fire/internal/applications"
)

func (a *API) GetDeploymentById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "deploymentid")

	deployment, err := a.deployments.GetDeploymentById(context.Background(), id)

	if err != nil {
		if err == applications.ErrApplicationNotFound {
			slog.Error("ErrApplicationNotFound", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if err == sql.ErrNoRows {
			slog.Error("No deployment defined yet", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		slog.Error("error while getting deployment", err)
		InternalServerError(w, err)
		return
	}

	AnswerJson(w, deployment, http.StatusOK)
}
