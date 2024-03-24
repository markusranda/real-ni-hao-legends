package game

import (
	"math/rand"
	"ni-hao-legends/game/buildings"
	"ni-hao-legends/models"
)

var possibilities = []func(stats buildings.BuildingStats) models.NHBuilding{
	buildings.CreateDreamPavilion,
	buildings.CreateHoleInTheGround,
	buildings.CreateNMSGjenbruk,
	buildings.CreateOrientalDragon,
	buildings.CreateEricClapton,
	buildings.CreateØysteinSunde,
	buildings.CreateFruitCasket,
	buildings.CreateKjellElvis,
}

func GetRandomLoot() models.NHBuilding {
	randomIndex := rand.Intn(len(possibilities))
	// Get a random building from the possibilities array
	return possibilities[randomIndex](buildings.RandomStats)
}
