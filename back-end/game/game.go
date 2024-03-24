package game

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"ni-hao-legends/database"
	"ni-hao-legends/models"
)

var State models.GameState

var handlers = map[string]func(models.Command) error{
	"init":                       Init,
	"chat":                       Chat,
	"building.upgrade":           BuildingUpgrade,
	"building.move_to_town":      BuildingMoveToTown,
	"building.move_to_inventory": BuildingMoveToInventory,
	"samba.score":                ScoreUpdate,
	"cheats.loot":                Loot,
}

func InitGameState() {
	// Query the states table
	var stateData string
	err := database.DB.QueryRow("SELECT state_data FROM states WHERE id = 1").Scan(&stateData)
	if err != nil {
		if err == sql.ErrNoRows {
			// If no row found, insert the row
			defaultStateData := models.GameState{
				Players:  make(map[models.PlayerId]*models.PlayerGameState),
				Chat:     models.Chat{Messages: make([]models.ChatMessage, 0)},
				Samba:    models.NHSamba{Scores: make([]models.NHScore, 0)},
				Commands: make([]models.Command, 0),
			}
			defaultStateDataStr, jsonErr := json.Marshal(defaultStateData)
			if jsonErr != nil {
				panic(fmt.Sprintf("Failed to marshall the games state.. %s\n", jsonErr.Error()))
			}

			_, err = database.DB.Exec("INSERT INTO states(id, state_data) VALUES (1, ?)", defaultStateDataStr)
			if err != nil {
				fmt.Println("Error inserting row:", err)
				return
			}
			fmt.Println("Inserted default row into states table.")
			State = defaultStateData
		} else {
			fmt.Println("Error querying database:", err)
			return
		}
	}

	err = json.Unmarshal([]byte(stateData), &State)
	if err != nil {
		log.Fatal("Error unmarshalling state data:", err)
	}
}

func CloseGameState() {
	// Serialize the GameState struct into JSON
	stateData, err := json.Marshal(State)
	if err != nil {
		fmt.Printf("error marshalling game state: %v\n", err)
	}

	// Update the states table with the serialized state data
	_, err = database.DB.Exec("UPDATE states SET state_data = ? WHERE id = 1", stateData)
	if err != nil {
		fmt.Printf("error updating game state in database: %v\n", err)
	}
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
