package api

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *API) listDeployment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	apps, err := a.deployments.ListApplicationDeployments(context.Background(), id)
	if err != nil {
		if err == errors.New("application not found") {
			slog.Error("ErrApplicationNotFound", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		slog.Error("error listing application", err)
		return
	}

	AnswerJson(w, apps, 200)
}
