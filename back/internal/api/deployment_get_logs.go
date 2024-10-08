package api

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

func (a *API) getLogs(w http.ResponseWriter, r *http.Request) {
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
					log.Printf("logs WebSocket closed unexpectedly : %v", err)
				} else {
					log.Println("logs WebSocket closed proprely :", err)
				}
				break
			}
		}
	}()
	var previousLogs string = ""

	for {
		containerLogs, err := a.deployments.GetContainerLogs(context.Background(), app_id)
		if err != nil {
			log.Println("error while getting container stats: %w", err)
		}
		if previousLogs != containerLogs {
			previousLogs = containerLogs
			if err := conn.WriteMessage(websocket.TextMessage, []byte(containerLogs)); err != nil {
				log.Println("error while sending  :", err)
				break
			}
		}
	}
}
