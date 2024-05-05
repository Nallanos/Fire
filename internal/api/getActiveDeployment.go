package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nallanos/fire/internal/applications"
)

func (a *API) getActiveDeployment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	deployment, err := a.deployments.GetLatestDeployment(context.Background(), id)

	if err != nil {
		if err == applications.ErrApplicationNotFound {
			slog.Error("ErrApplicationNotFound", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		slog.Error("error while getting deployment", err)
		InternalServerError(w, err)
		return
	}

	AnswerJson(w, deployment, http.StatusOK)
}
