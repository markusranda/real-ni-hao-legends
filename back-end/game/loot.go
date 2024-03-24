package game

import (
	"math/rand"
	"ni-hao-legends/game/buildings"
	"ni-hao-legends/models"
	"sort"
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

var possibleBuildings []models.NHBuilding
var weights []float64
var totalWeight float64

func init() {
	// Create one of each possibility
	for _, possibility := range possibilities {
		possibleBuildings = append(possibleBuildings, possibility(buildings.RandomStats))
	}

	/* Calculate the weights
	Common (Rarity = 0): The weight is 1.0 / (0 + 1) = 1.0
	Uncommon (Rarity = 1): The weight is 1.0 / (1 + 1) = 0.5
	Rare (Rarity = 2): The weight is 1.0 / (2 + 1) = 0.33
	Epic (Rarity = 3): The weight is 1.0 / (3 + 1) = 0.25
	Legendary (Rarity = 4): The weight is 1.0 / (4 + 1) = 0.2
	The total weight is the sum of all weights, which is 1.0 + 0.5 + 0.33 + 0.25 + 0.2 = 2.28  So, the chances of getting each rarity level are:
	Common: 1.0 / 2.28 = 0.44 or 44%
	Uncommon: 0.5 / 2.28 = 0.22 or 22%
	Rare: 0.33 / 2.28 = 0.14 or 14%
	Epic: 0.25 / 2.28 = 0.11 or 11%
	Legendary: 0.2 / 2.28 = 0.09 or 9%
	*/
	weights = make([]float64, len(possibilities))
	for i, building := range possibleBuildings {
		// Higher rarity level means lower chance of being selected
		weights[i] = 1.0 / float64(building.Rarity+1)
		if i > 0 {
			weights[i] += weights[i-1]
		}
	}
	totalWeight = weights[len(weights)-1]
}

func GetRandomLoot() models.NHBuilding {
	randomValue := rand.Float64() * totalWeight
	index := sort.Search(len(weights), func(i int) bool { return weights[i] >= randomValue })
	return possibleBuildings[index]
}
