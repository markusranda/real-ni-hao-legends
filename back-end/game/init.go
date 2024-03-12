package game

import (
	"log"
	"ni-hao-legends/game/buildings"
	"ni-hao-legends/models"
)

func Init(command models.Command) {
	/*
		Add player to game state
	*/
	playerId := command.PlayerId
	log.Printf("Init player %s", playerId)

	if State.Players[playerId] == nil {
		State.Players[playerId] = &models.PlayerGameState{
			Inventory:      models.Inventory{Buildings: []models.NHBuilding{}},
			RetirementFund: 0.0,
			Town: models.NHTown{
				Money: 0.0,
				Name:  command.Data["name"].(string),
				Buildings: map[string]models.NHBuilding{
					"Dream pavilion": buildings.CreateDreamPavilion(buildings.InitStats),
				},
			},
		}
	}
}
