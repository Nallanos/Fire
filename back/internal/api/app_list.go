package api

import (
	"log/slog"
	"net/http"
)

func (a *API) listApps(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(UserKey)
	apps, err := a.apps.ListApplications(r.Context(), userID.(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		slog.Error(err.Error())
		return
	}

	AnswerJson(w, apps, 200)
}
