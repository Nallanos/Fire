package api

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *API) startContainer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := a.deployments.StartContainer(id)
	if err != nil {
		slog.Error("error starting container", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		InternalServerError(w, err)
		return
	}
}
