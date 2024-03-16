package api

import (
	"database/sql"
	"log/slog"

	"net/http"

	"github.com/docker/docker/client"
	_ "github.com/mattn/go-sqlite3"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/nallanos/fire/internal/applications"
	"github.com/nallanos/fire/internal/db"
)

type API struct {
	router *chi.Mux
	apps   *applications.Service
	docker *client.Client
}

func NewAPI(docker *client.Client) (*API, error) {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	sqliteDb, err := sql.Open("sqlite3", "./db/database.sqlite3")
	if err != nil {
		slog.Error("Error opening database", err)
		return nil, err
	}

	appsService := applications.NewService(*db.New(sqliteDb))

	api := &API{
		apps:   appsService,
		router: r,
		docker: docker,
	}

	r.Get("/health", api.health)
	r.Post("/apps", api.createApp)
	r.Get("/apps", api.listApps)
	r.Get("/apps/{id}", api.getApp)
	r.Delete("/apps/{id}", api.deleteApp)
	return api, nil
}

func (api *API) ListenAndServe() error {
	slog.Info("Listening on :8080")
	return http.ListenAndServe(":8080", api.router)
}
