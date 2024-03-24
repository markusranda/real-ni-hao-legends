package game

import (
	"fmt"
	"log"
	"ni-hao-legends/models"
)

var State = models.GameState{
	Players:  make(map[models.PlayerId]*models.PlayerGameState),
	Chat:     models.Chat{Messages: make([]models.ChatMessage, 0)},
	Samba:    models.NHSamba{Scores: make([]models.NHScore, 0)},
	Commands: make([]models.Command, 0),
}

var handlers = map[string]func(models.Command) error{
	"init":                       Init,
	"chat":                       Chat,
	"building.upgrade":           BuildingUpgrade,
	"building.move_to_town":      BuildingMoveToTown,
	"building.move_to_inventory": BuildingMoveToInventory,
	"samba.score":                ScoreUpdate,
	"cheats.loot":                Loot,
	"effect.nuke":                Nuke,
	"effect.roll_ball":           RollBall,
}

func HandleCommand(command models.Command) error {
	log.Printf("Handling command: %+v", command)
	handler, ok := handlers[command.Action]
	if !ok {
		return fmt.Errorf("unknown command action: %s", command.Action)
	}

	State.Commands = append(State.Commands, command)
	// Limit the number of commands to 10
	if len(State.Commands) > 10 {
		// Remove the oldest command
		State.Commands = State.Commands[1:]
	}

	return handler(command)
}
