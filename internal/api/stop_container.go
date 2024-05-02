package api

import (
	"context"
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

	err = a.deployments.StopContainer(app)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		InternalServerError(w, err)
		return
	}
}
