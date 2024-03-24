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
	go func() {
		for {
			delay := time.Duration(5) * time.Second
			time.Sleep(delay)
			for _, npc := range npcs {

				SendCommandOverWs(models.Command{
					PlayerId: npc.id,
					Action:   "building.upgrade",
					Data:     map[string]interface{}{"buildingId": "Dream pavilion"},
				})

			}
		}
	}()

	// random tick to chat
	go func() {
		for {
			state := &game.State
			delay := time.Duration(rand.Intn(180)+1) * time.Second
			time.Sleep(delay)

			//pick a random npc
			npc := npcs[rand.Intn(len(npcs))]

			SendCommandOverWs(models.Command{
				PlayerId: npc.id,
				Action:   "chat",
				Data: map[string]interface{}{
					"message":  generateChatMessage(state, npc.id),
					"type":     "",
					"userId":   npc.id,
					"userName": npc.name,
				},
			})
		}
	}()
}

func generateChatMessage(gs *models.GameState, playerId string) string {

	selfTalkMessages := []string{
		"thinking about why coffee smells better than it tastes ☕️🤔",
		"if i were a cat, would i enjoy human videos? 🐱💻",
		"do fish ever get thirsty, or is that just their life? 🐠💧",
		"why do we say 'heads up' when we actually duck? 🤷‍♂️",
		"pondering if my plants are judging my music taste 🌱🎶",
		"sometimes i wonder if my fridge lights up when i'm not around 🚪💡",
		"counting the number of steps to the fridge as exercise today 🚶‍♂️🍔",
	}
	rand.Seed(time.Now().UnixNano())

	var players []string
	for id := range gs.Players {
		if id != playerId {
			players = append(players, id)
		}
	}

	var message string
	// choose a random player to trash talk
	if len(players) > 0 {
		otherPlayerId := players[rand.Intn(len(players))]
		otherPlayerName := gs.Players[otherPlayerId].Town.Name

		// decide whether to trash talk or self talk
		if rand.Intn(2) == 0 {
			message = selfTalkMessages[rand.Intn(len(selfTalkMessages))]
			return message
		}

		otherPlayerMessages := []string{
			fmt.Sprintf("'%s? more like “try-hard town” 😂 get on my level.'", otherPlayerName),
			fmt.Sprintf("'heard %s was planning something big. too bad, i’m already three steps ahead 🚀'", otherPlayerName),
			fmt.Sprintf("'lol, just outperformed %s with my eyes closed. EZ game EZ life 🕶️'", otherPlayerName),
			fmt.Sprintf("'%s thinks they're a strategist? watched their play, yawned, then did it better 🧠💤'", otherPlayerName),
			fmt.Sprintf("'%s's latest move? cute. but i’m playing 4D chess over here 🐱‍👤'", otherPlayerName),
			fmt.Sprintf("'%s bragging again? wake me up when they actually win something 😴💤'", otherPlayerName),
			fmt.Sprintf("'%s's effort? admirable but as effective as a screen door on a submarine 🚢🚪'", otherPlayerName),
			fmt.Sprintf("'%s? more like precious pie 🍰! let’s have a sweet game together!'", otherPlayerName),
			fmt.Sprintf("'aww, %s's trying their best! let's cheer them on 📣! everyone's a star 🌟 in ni hao radio universe!'", otherPlayerName),
			fmt.Sprintf("'%s, your strategies are as cute as a button! can't wait to see what you do next 🌈✨'", otherPlayerName),
			fmt.Sprintf("'%s, watching your plays is like watching fireworks! beautiful but i won’t lose! 🎆🤗'", otherPlayerName),
			fmt.Sprintf("'%s's got spirit! let's show them our moves too, but with extra glitter! ✨💃'", otherPlayerName),
			fmt.Sprintf("'%s, your town’s got charm! let's make this match a memorable one 💌💖'", otherPlayerName),
			fmt.Sprintf("'isn't %s just adorable trying so hard? let's give them a round of applause 👏🐾! together, we shine brighter!'", otherPlayerName),
			fmt.Sprintf("'%s? thought that was a new pizza topping 🍕'", otherPlayerName),
			fmt.Sprintf("'%s’s strategy? thought it was a recipe for pancakes 🥞'", otherPlayerName),
			fmt.Sprintf("'watched %s play. now i want a pet duck 🦆'", otherPlayerName),
			fmt.Sprintf("'%s? isn't that the latest dance move? 🕺'", otherPlayerName),
			fmt.Sprintf("'just saw %s’s move. did a magic trick instead. 🎩✨'", otherPlayerName),
			fmt.Sprintf("'%s bragging? thought they found the last potato chip 🍟'", otherPlayerName),
			fmt.Sprintf("'meeting %s in-game was like finding socks in the fridge 🧦❄️'", otherPlayerName),
		}

		message = otherPlayerMessages[rand.Intn(len(otherPlayerMessages))]
	}

	return message
}
