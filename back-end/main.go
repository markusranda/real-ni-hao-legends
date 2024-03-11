package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"math"
	"net/http"
	"ni-hao-legends/game"
	"ni-hao-legends/models"
	"ni-hao-legends/npc"
	"sync"
	"time"
)

type Connection struct {
	ws   *websocket.Conn
	send chan []byte // Buffered channel of outbound messages.
}

type Hub struct {
	connections map[*Connection]bool
	broadcast   chan []byte
	register    chan *Connection
	unregister  chan *Connection
	mutex       sync.Mutex
}

var hub = Hub{
	broadcast:   make(chan []byte),
	register:    make(chan *Connection),
	unregister:  make(chan *Connection),
	connections: make(map[*Connection]bool),
}

func (h *Hub) run() {
	for {
		select {
		case conn := <-h.register:
			h.mutex.Lock()
			h.connections[conn] = true
			h.mutex.Unlock()
		case conn := <-h.unregister:
			if _, ok := h.connections[conn]; ok {
				h.mutex.Lock()
				delete(h.connections, conn)
				close(conn.send)
				h.mutex.Unlock()
			}
		case message := <-h.broadcast:
			h.mutex.Lock()
			for conn := range h.connections {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(h.connections, conn)
				}
			}
			h.mutex.Unlock()
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	conn := &Connection{send: make(chan []byte, 256), ws: ws}
	hub.register <- conn

	go func() {
		for {
			select {
			case message, ok := <-conn.send:
				if !ok {
					err := ws.WriteMessage(websocket.CloseMessage, []byte{})
					if err != nil {
						log.Printf("❌️🥵 ws.WriteMessage error: %v", err)
					}
					return
				}
				err := ws.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Printf("❌️🥵 ws.WriteMessage error: %v", err)
				}
			}
		}
	}()

	defer func() { hub.unregister <- conn }()

	for {
		_, message, err := ws.ReadMessage()
		log.Println("⬅️ request", string(message))
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("❌️🥵 error: %v", err)
			}
			break
		}
		var cmd models.Command

		err = json.Unmarshal(message, &cmd)
		log.Println(cmd.Time)
		cmd.Time = time.Now().Unix()
		if err != nil {
			log.Println("❌️🥵 error: %v", err)
			continue
		}

		if cmd.Action == "" {
			log.Println("❌️🥵 hold up there pall: !!!! empty type !!!!")
			continue
		}

		// Updates the game state
		err = game.HandleCommand(cmd)

		if err != nil {
			log.Println("❌️🥵error handling command: %v", err)
			continue
		}

		BroadcastGameState()
	}
}

func BroadcastGameState() {
	responseJson, err := json.Marshal(game.State)
	if err != nil {
		log.Printf("❌️🥵 broadcastGameState: error marshalling json: %v", err)
		return
	}
	hub.broadcast <- responseJson
}

// function that updates the money of each connected player every second
func lønningsdag() {
	log.Printf("🕒 starting lønningsdag")
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		log.Printf("🕒 update money")
		select {
		case <-ticker.C:
			for _, playerState := range game.State.Players {
				totalIncome := 0.0
				for _, building := range playerState.Town.Buildings {
					totalIncome += building.Income
				}
				newMoney := math.Floor(playerState.Town.Money + totalIncome)
				playerState.Town.Money = newMoney
			}
			BroadcastGameState()
		}
	}
}

func main() {
	go hub.run()
	http.HandleFunc("/ws", wsHandler)
	log.Println("✨ Server started at :8080 ✨")
	go lønningsdag()

	go func() {
		time.Sleep(5 * time.Second)
		npc.InitNpcBehaviour()
		log.Println("🤖 NPCs started")
	}()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
