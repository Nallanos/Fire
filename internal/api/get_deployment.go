package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nallanos/fire/internal/applications"
)

func (a *API) getDeployment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	deployments, err := a.deployments.GetDeployment(id, context.Background())

	if err != nil {
		if err == applications.ErrApplicationNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		InternalServerError(w, err)
		return
	}
	AnswerJson(w, deployments, http.StatusOK)
}
