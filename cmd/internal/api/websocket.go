package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleAlerts(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	for {
		// Read message
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		// Broadcast alert (in real app, would use Redis Pub/Sub)
		log.Printf("Received alert: %s", msg)
	}
}
