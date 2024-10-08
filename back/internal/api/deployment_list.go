package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *API) listDeployment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	apps, err := a.deployments.ListApplicationDeployments(context.Background(), id)
	if err != nil {
		if err.Error() == "error getting deployment: sql: no rows in result set" {
			AnswerJson(w, nil, http.StatusOK)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		slog.Error(err.Error())
		return
	}
	AnswerJson(w, apps, 200)
}
