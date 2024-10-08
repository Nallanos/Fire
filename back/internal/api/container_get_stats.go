package api

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (a *API) getStatsOfActiveDeployment(w http.ResponseWriter, r *http.Request) {
	app_id := chi.URLParam(r, "id")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erreur lors de l'upgrade :", err)
		return
	}
	defer conn.Close()
	go func() {
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				if err.(*websocket.CloseError).Code == websocket.CloseGoingAway || err.(*websocket.CloseError).Code == websocket.CloseAbnormalClosure {
					log.Printf("stats WebSocket closed unexpectedly : %v", err)
				} else {
					log.Println("stats WebSocket closed proprely :", err)
				}
				break
			}
		}
	}()
	for {
		containerStats, err := a.deployments.GetContainerStats(context.Background(), app_id)
		if err != nil {
			log.Println("error while getting container stats: %w", err)
		}
		if err := conn.WriteJSON(containerStats); err != nil {
			log.Println("error while sending  :", err)
			break
		}
	}
}
