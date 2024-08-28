package api

import (
	"log/slog"
	"net/http"
)

func (a *API) listApps(w http.ResponseWriter, r *http.Request) {
	apps, err := a.apps.ListApplications(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		slog.Error(err.Error())
		return
	}

	AnswerJson(w, apps, 200)
}
