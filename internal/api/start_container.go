package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *API) startContainer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	app, err := a.apps.GetApplication(context.Background(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		InternalServerError(w, err)
		return
	}

	err = a.deployments.StartContainer(app)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		InternalServerError(w, err)
		return
	}
	AnswerJson(w, nil, http.StatusOK)
}
