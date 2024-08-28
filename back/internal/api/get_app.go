package api

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nallanos/fire/internal/applications"
)

func (a *API) getApp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	app, err := a.apps.GetApplication(r.Context(), id)
	if err != nil {
		if err == applications.ErrApplicationNotFound {
			http.Error(w, "Application not found", http.StatusNotFound)
			return
		}
		slog.Error(err.Error())
		InternalServerError(w, err)
		return
	}

	AnswerJson(w, app, http.StatusOK)
}
