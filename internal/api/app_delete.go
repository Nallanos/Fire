package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *API) deleteApp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := a.apps.DeleteApplication(r.Context(), id)
	if err != nil {
		http.Error(w, "error while deleting application", http.StatusNotFound)
		InternalServerError(w, err)
		return
	}
	AnswerJson(w, nil, http.StatusOK)
}
