package game

import (
	"log"
	"ni-hao-legends/models"
)

func BuildingUpgrade(command models.Command) {
	id, ok := command.Data["buildingId"].(string)
	if !ok {
		log.Printf("buildingId is not a string")
	}
	playerId := command.PlayerId
	state := State.Players[playerId]
	building := state.Town.Buildings[id]

	if state.Town.Money < (building.UpgradeCost) {
		log.Printf("❌💰 not enough money")
		return
	}

	building.Upgrade()

	state.Town.Buildings[id] = building
}
