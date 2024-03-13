package npc

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"ni-hao-legends/models"
	"time"
)

var npcConnections = make(map[string]*websocket.Conn)

func ConnectNpcToWs(npc npc) {
	url := "ws://localhost:8080/ws"

	var c *websocket.Conn
	var err error

	for i := 0; i < 5; i++ {
		c, _, err = websocket.DefaultDialer.Dial(url, nil)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to %s. Retrying in 5 seconds...", url)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatal("dial:", err)
	}

	npcConnections[npc.id] = c
}

func SendCommandOverWs(command models.Command) {
	npcId := command.PlayerId
	conn := npcConnections[npcId]
	if conn == nil {
		log.Printf("Connection for NPC %s not found", npcId)
		return
	}
	message, _ := json.Marshal(command)
	conn.WriteMessage(websocket.TextMessage, message)
}
