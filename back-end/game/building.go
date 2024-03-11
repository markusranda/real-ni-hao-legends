package game

import (
	"log"
	"math"
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

	if state.Town.Money < (building.BaseCost * building.Level) {
		log.Printf("❌💰 not enough money")
		return
	}

	upgradeCost := math.Floor(building.BaseCost * building.Level * 1.1)
	building.UpgradeCost = upgradeCost
	state.Town.Money -= upgradeCost
	building.Level++
	building.Income = math.Floor(building.Income * 1.1)

	state.Town.Buildings[id] = building
}
