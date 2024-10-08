package api

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nallanos/fire/internal/applications"
)

func (a *API) getDeploymentById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "deploymentid")

	deployment, err := a.deployments.GetDeploymentById(context.Background(), id)

	if err != nil {
		if err == applications.ErrApplicationNotFound {
			slog.Error(fmt.Sprintf("ErrApplicationNotFound %v", err))
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if err == sql.ErrNoRows {
			slog.Error(fmt.Sprintf("No deployment defined yet %v", err))
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		slog.Error(err.Error())
		InternalServerError(w, err)
		return
	}
	AnswerJson(w, deployment, http.StatusOK)
}
