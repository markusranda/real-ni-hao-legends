package game

import (
	"ni-hao-legends/models"
)

func Loot(command models.Command) {
	userId := command.PlayerId
	State.Players[userId].Inventory.Buildings = append(State.Players[userId].Inventory.Buildings, GetRandomLoot())
}
