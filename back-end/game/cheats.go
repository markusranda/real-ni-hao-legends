package game

import (
	"ni-hao-legends/models"
)

func Loot(command models.Command) error {
	userId := command.PlayerId
	State.Players[userId].Inventory.Buildings = append(State.Players[userId].Inventory.Buildings, GetRandomLoot())

	return nil
}
