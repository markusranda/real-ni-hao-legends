package npc

import (
	"fmt"
	"math/rand"
	"ni-hao-legends/game"
	"ni-hao-legends/models"
	"time"
)

type npc struct {
	id   string
	name string
}

func InitNpcBehaviour() {
	npcs := []npc{
		{
			id:   "npc1",
			name: "Vodcouch City",
		},
		{
			id:   "npc2",
			name: "Playa de l'annee",
		},
	}

	for _, npcId := range npcs {
		ConnectNpcToWs(npcId)
	}

	// INIT
	for _, npc := range npcs {
		SendCommandOverWs(models.Command{
			PlayerId: npc.id,
			Action:   "init",
			Data: map[string]interface{}{
				"name": npc.name,
			},
		})
	}

	// constant tick to do stuff
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	go func() {
		for {
			select {
			case <-ticker.C:
				for _, npc := range npcs {
					SendCommandOverWs(models.Command{
						PlayerId: npc.id,
						Action:   "building.upgrade",
						Data:     map[string]interface{}{"buildingId": "Dream pavilion"},
					})
				}

			}
		}
	}()

	// random tick to chat
	go func() {
		for {
			state := &game.State
			delay := time.Duration(rand.Intn(30)+1) * time.Second
			time.Sleep(delay)

			SendCommandOverWs(models.Command{
				PlayerId: "npc1",
				Action:   "chat",
				Data: map[string]interface{}{
					"message":  generateChatMessage(state, "npc1"),
					"type":     "",
					"userId":   "npc1",
					"userName": "npc1",
				},
			})

			SendCommandOverWs(models.Command{
				PlayerId: "npc2",
				Action:   "chat",
				Data: map[string]interface{}{
					"message":  generateChatMessage(state, "npc2"),
					"type":     "",
					"userId":   "npc2",
					"userName": "npc2",
				},
			})
		}
	}()
}

func generateChatMessage(gs *models.GameState, playerId models.PlayerId) string {
	rand.Seed(time.Now().UnixNano())

	players := make([]models.PlayerId, 0, len(gs.Players))
	for id := range gs.Players {
		if id != playerId {
			players = append(players, id)
		}
	}

	var trashTalks []string
	if len(players) > 0 {
		for _, otherPlayerId := range players {
			otherPlayerName := gs.Players[otherPlayerId].Town.Name

			otherPlayerMessages := []string{
				fmt.Sprintf("'%s? More like amateur hour. Watch and learn.'", otherPlayerName),
				fmt.Sprintf("'%s's got plans? Mine are already in motion. Step aside.'", otherPlayerName),
				fmt.Sprintf("'Just outdid %s's best in half the time. EZ.'", otherPlayerName),
				fmt.Sprintf("'%s's strategy? Saw it, laughed, improved it.'", otherPlayerName),
				fmt.Sprintf("'%s bragging again? Wake me up when there's real competition.'", otherPlayerName),
			}
			trashTalks = append(trashTalks, otherPlayerMessages[rand.Intn(len(otherPlayerMessages))])
		}
	}

	var message string
	if len(trashTalks) > 0 {
		message = trashTalks[rand.Intn(len(trashTalks))]
	} else {
		message = "No trash talk available"
	}

	return message
}
