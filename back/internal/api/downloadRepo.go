package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *API) downloadRepo(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "id")

	fmt.Println(id, "repo downloaded")
}
