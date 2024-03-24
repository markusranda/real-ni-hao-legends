package game

import (
	"fmt"
	"github.com/google/uuid"
	"math"
	"ni-hao-legends/models"
	"time"
)

func Nuke(command models.Command) error {
	// Nuke a player's town
	target, ok := command.Data["target"].(string)
	if !ok {
		return fmt.Errorf("nuke: no target specified")
	}

	// todo use playerid
	var player *models.PlayerGameState = nil
	for _, p := range State.Players {
		if p.Town.Name == target {
			player = p
			break
		}
	}

	if player == nil {
		return fmt.Errorf("nuke: target not found")
	}

	player.Town.Money = math.Floor(player.Town.Money * .1)
	nukeEffect := models.NHEffect{
		EffectType: "nuke",
		Timestamp:  command.Time,
		UUID:       uuid.New(),
	}
	player.Effects = append(player.Effects, nukeEffect)

	go func() {
		time.Sleep(10 * time.Second)
		for i, effect := range player.Effects {
			if effect.UUID == nukeEffect.UUID {
				// Remove the nuke effect from the player's effects
				player.Effects = append(player.Effects[:i], player.Effects[i+1:]...)
				break
			}
		}
	}()

	return nil
}
