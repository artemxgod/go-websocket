package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/ws", handleConnections)
	log.Println("started server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	ws.SetReadDeadline(time.Now().Add(60 * time.Second))

	for {
		messageType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("received message: ", msg)

		if err := ws.WriteMessage(messageType, msg); err != nil {
			log.Println(err)
			break
		}

		ws.SetReadDeadline(time.Now().Add(60 * time.Second))
	}
}
