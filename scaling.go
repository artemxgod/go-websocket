package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	go handleConnectionAsync(conn)
}

func handleConnection(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println(p, messageType)
		// обработка сообщения...
	}
}

func handleConnectionAsync(conn *websocket.Conn) {
	msgChan := make(chan []byte)

	go func() {
		for {
			message, ok := <-msgChan
			if !ok {
				return
			}
			conn.WriteMessage(websocket.TextMessage, message)
		}
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			close(msgChan)
			break
		}
		msgChan <- message
	}
}
