package api

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nallanos/fire/internal/applications"
)

func (a *API) deployApp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	app, err := a.apps.GetApplication(r.Context(), id)
	if err != nil {
		if err == applications.ErrApplicationNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		InternalServerError(w, err)
		return
	}

	err = a.deployments.Deploy(app)

	if err != nil {
		if err == applications.ErrApplicationNotFound {
			slog.Error("ErrApplicationNotFound", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		slog.Error("error while depoying app", err)
		InternalServerError(w, err)
		return
	}
}
