package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nallanos/fire/internal/applications"
)

func (a *API) UpdateApplicationStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := a.apps.UpdateAppStatus(context.Background(), id)

	if err != nil {
		if err == applications.ErrApplicationNotFound {
			slog.Error("Error updating application status", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		return
	}
}
