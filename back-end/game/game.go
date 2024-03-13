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
}

func HandleCommandV2(command models.Command) error {
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

func HandleCommand(command models.Command) error {
	log.Printf("Handling command: %+v", command)
	switch command.Action {
	case "init":
		Init(command)
	case "chat":
		Chat(command)
	case "building.upgrade":
		err := BuildingUpgrade(command)
		if err != nil {
			return err
		}
	case "building.move_to_town":
		err := BuildingMoveToTown(command)
		if err != nil {
			return err
		}
	case "building.move_to_inventory":
		err := BuildingMoveToInventory(command)
		if err != nil {
			return err
		}
	case "samba.score":
		ScoreUpdate(command)
	case "cheats.loot":
		Loot(command)
	default:
		return fmt.Errorf("unknown command action: %s", command.Action)
	}
	return nil
}
