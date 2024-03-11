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

func HandleCommand(command models.Command) error {
	log.Printf("Handling command: %+v", command)
	switch command.Action {
	case "init":
		Init(command)
	case "chat":
		Chat(command)
	case "building.upgrade":
		BuildingUpgrade(command)
	case "samba.score":
		ScoreUpdate(command)
	default:
		return fmt.Errorf("unknown command action: %s", command.Action)
	}

	State.Commands = append(State.Commands, command)
	// Limit the number of commands to 10
	if len(State.Commands) > 10 {
		// Remove the oldest command
		State.Commands = State.Commands[1:]
	}

	return nil
}
