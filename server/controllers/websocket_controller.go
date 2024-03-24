package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"githut.com/goredan/server/models"
	"log"
	"net/http"
)

type WebSocketController struct{}

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan models.ChatMessage)
	upgrader  = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func readMessages(ws *websocket.Conn) {
	for {
		var msg models.ChatMessage
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		log.Println("Receive:", msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		log.Println("Sending:", msg)
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func (home *WebSocketController) HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Close connection when everything done
	defer ws.Close()

	// Register new client
	clients[ws] = true

	// Read messages
	readMessages(ws)

	// Grab messages and send to all clients
	go handleMessages()
}
