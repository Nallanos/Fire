package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *API) getActiveDeployment(w http.ResponseWriter, r *http.Request) {
	app_id := chi.URLParam(r, "id")
	deployment, err := a.deployments.GetLatestDeployment(context.Background(), app_id)

	if err != nil {
		if err.Error() == "error getting deployment: sql: no rows in result set" {
			AnswerJson(w, nil, http.StatusOK)
			return
		}
		slog.Error(err.Error())
		InternalServerError(w, err)
		return
	}

	AnswerJson(w, deployment, http.StatusOK)
}
