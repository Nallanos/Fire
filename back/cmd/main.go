package main

import (
	"log"
	"log/slog"

	"github.com/docker/docker/client"
	"github.com/nallanos/fire/internal/api"
)

func main() {
	client, err := client.NewClientWithOpts(client.WithHost("unix:///var/run/docker.sock"), client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal("Error creating docker client", err)
		panic(err)
	}

	defer client.Close()

	slog.Info("Starting Fire")

	slog.Info("Starting API")

	app, err := api.NewAPI(client)
	if err != nil {
		log.Fatal("Error creating API", err)
		panic(err)
	}

	err = app.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting API", err)
		panic(err)
	}
}
