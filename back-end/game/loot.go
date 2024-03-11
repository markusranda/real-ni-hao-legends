package game

import (
	"math/rand"
	"ni-hao-legends/game/buildings"
	"ni-hao-legends/models"
)

var possibilities []func(int) models.NHBuilding = []func(int) models.NHBuilding{
	buildings.CreateDreamPavilion,
	buildings.CreateHoleInTheGround,
	buildings.CreateNMSGjenbruk,
	buildings.CreateOrientalDragon,
}

func GetRandomLoot() models.NHBuilding {
	randomIndex := rand.Intn(len(possibilities))
	// Get a random building from the possibilities array
	return possibilities[randomIndex](buildings.RANDOM)
}
