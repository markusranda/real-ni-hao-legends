package game

import (
	"fmt"
	"github.com/google/uuid"
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

func BuildingMoveToTown(command models.Command) error {
	uniqueIdString := command.Data["uniqueId"].(string)
	uniqueId, err := uuid.Parse(uniqueIdString)
	if err != nil {
		panic("buildingUniqueKey is not an UUID")
	}
	playerId := command.PlayerId
	state := State.Players[playerId]

	var building *models.NHBuilding = nil
	for _, buildingInInventory := range state.Inventory.Buildings {
		if buildingInInventory.UniqueId == uniqueId {
			building = &buildingInInventory
			break
		}
	}

	if building == nil {
		panic("building not found in inventory")
	}

	_, buildingAlreadyExistsInTown := state.Town.Buildings[building.Key]
	if buildingAlreadyExistsInTown {
		return fmt.Errorf("building already in town")
	} else {
		state.Town.Buildings[building.Key] = *building
	}

	return nil
}

func BuildingMoveToInventory(command models.Command) error {
	id, ok := command.Data["buildingId"].(string)
	if !ok {
		panic("buildingId is not a string")
	}
	playerId := command.PlayerId
	state := State.Players[playerId]

	building, buildingExistsInTown := state.Town.Buildings[id]
	if !buildingExistsInTown {
		return fmt.Errorf("building not found in town")
	}

	state.Inventory.Buildings = append(state.Inventory.Buildings, building)
	delete(state.Town.Buildings, id)

	return nil
}
