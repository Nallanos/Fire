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
	"github.com/nallanos/fire/internal/deployment"
)

type API struct {
	router      *chi.Mux
	apps        *applications.Service
	docker      *client.Client
	deployments *deployment.ContainerService
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
	queries := db.New(sqliteDb)
	deploymentsService := deployment.NewService(docker, queries)
	appsService := applications.NewService(queries, deploymentsService)

	api := &API{
		apps:        appsService,
		router:      r,
		docker:      docker,
		deployments: deploymentsService,
	}

	r.Get("/health", api.health)

	r.Post("/apps", api.createApp)
	r.Get("/apps", api.listApps)

	r.Get("/apps/{id}", api.getApp)

	r.Get("/apps/{id}/deployment", api.listDeployment)
	r.Get("/apps/{id}/deployment/activeDeployment", api.getActiveDeployment)

	r.Post("/apps/{id}/deploy", api.deployApp)
	r.Post("/apps/{id}/stop", api.stopContainer)
	r.Post("/apps/{id}/start", api.startContainer)

	r.Post("/apps/{id}/download", api.downloadRepo)

	r.Get("/apps/{id}/getDeployment/{DeploymentId}", api.GetDeploymentById)

	r.Delete("/apps/{id}", api.deleteApp)
	return api, nil
}

func (api *API) ListenAndServe() error {
	slog.Info("Listening on :8081")
	return http.ListenAndServe(":8081", api.router)
}
