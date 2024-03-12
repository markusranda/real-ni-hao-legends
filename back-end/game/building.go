package game

import (
	"fmt"
	"ni-hao-legends/models"
)

func BuildingUpgrade(command models.Command) error {
	id, ok := command.Data["buildingId"].(string)
	if !ok {
		panic("buildingId is not a string")
	}
	playerId := command.PlayerId
	state := State.Players[playerId]
	building := state.Town.Buildings[id]

	if state.Town.Money < (building.UpgradeCost) {
		return fmt.Errorf("not enough money to upgrade building")
	}

	building.Upgrade()

	state.Town.Buildings[id] = building
	return nil
}
