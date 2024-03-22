package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nallanos/fire/internal/applications"
)

func (a *API) deployApp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	app, err := a.apps.GetApplication(context.Background(), id)
	if err != nil {
		if err == applications.ErrApplicationNotFound {
			http.Error(w, "Application not found", http.StatusNotFound)
			return
		}

		InternalServerError(w, err)
		return
	}

	a.deployments.DeployApplication(app)

	err = a.deployments.DeployApplication(app)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	AnswerJson(w, app, http.StatusCreated)
}
