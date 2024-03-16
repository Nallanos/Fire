package api

import (
	"context"
	"net/http"

	"github.com/nallanos/fire/internal/applications"
)

type CreateAppBody struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Port  int    `json:"port"`
}

func (a *API) createApp(w http.ResponseWriter, r *http.Request) {
	var body applications.CreateApplicationOptions
	if err := DecodeJSON(r, &body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	app, err := a.apps.CreateApplication(context.Background(), body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		InternalServerError(w, err)
		return
	}
	AnswerJson(w, app, http.StatusCreated)
}
