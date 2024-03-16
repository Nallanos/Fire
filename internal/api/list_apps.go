package api

import (
	"net/http"
)

func (a *API) listApps(w http.ResponseWriter, r *http.Request) {
	apps, err := a.apps.ListApplications(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	AnswerJson(w, apps, 200)
}
