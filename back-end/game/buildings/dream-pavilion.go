package buildings

import (
	"ni-hao-legends/models"
)

func NewBuilding(key string, name string, imgUrl string, income float64, upgradeCost float64, incomeScale float64, upgradeCostScale float64, stats BuildingStats) models.NHBuilding {
	building := models.NHBuilding{
		Key:              key,
		Name:             name,
		ImgUrl:           imgUrl,
		Income:           income,
		Level:            1,
		UpgradeCost:      upgradeCost,
		IncomeScale:      incomeScale,
		UpgradeCostScale: upgradeCostScale,
	}

	return *maybeRandomizeStats(&building, stats)
}

func CreateDreamPavilion(
	stats BuildingStats,
) models.NHBuilding {
	return NewBuilding(
		"Dream pavilion",
		"Dream pavilion",
		"dream_pavilion.webp",
		10,
		1000,
		1.1,
		1.1,
		stats,
	)
}

func CreateHoleInTheGround(stats BuildingStats) models.NHBuilding {
	building := models.NHBuilding{
		Key:         "Hole in the ground",
		Name:        "Hole in the ground",
		ImgUrl:      "hole_in_the_ground.webp",
		Level:       1,
		Income:      -20,
		UpgradeCost: 3000,

		IncomeScale:      1.1,
		UpgradeCostScale: 1.1,
	}
	return *maybeRandomizeStats(&building, stats)
}

func CreateNMSGjenbruk(stats BuildingStats) models.NHBuilding {
	building := models.NHBuilding{
		Key:         "NMS Gjenbruk",
		Name:        "NMS Gjenbruk",
		ImgUrl:      "nms_gjenbruk.webp",
		Level:       1,
		Income:      10,
		UpgradeCost: 3000,

		IncomeScale:      1.1,
		UpgradeCostScale: 1.1,
	}
	return *maybeRandomizeStats(&building, stats)
}

func CreateOrientalDragon(stats BuildingStats) models.NHBuilding {
	building := models.NHBuilding{
		Key:         "Oriental Dragon",
		Name:        "Oriental Dragon",
		ImgUrl:      "oriental_dragon.webp",
		Income:      10,
		Level:       1,
		UpgradeCost: 4000000,

		IncomeScale:      1.5,
		UpgradeCostScale: 1.1,
	}
	return *maybeRandomizeStats(&building, stats)
}

func CreateEricClapton(stats BuildingStats) models.NHBuilding {
	building := models.NHBuilding{
		Key:         "Eric Clapton",
		Name:        "Eric Clapton",
		ImgUrl:      "eric_clapton.webp",
		Income:      20,
		Level:       1,
		UpgradeCost: 2000,

		IncomeScale:      1.1,
		UpgradeCostScale: 1.1,
	}

	return *maybeRandomizeStats(&building, stats)
}

func CreateØysteinSunde(stats BuildingStats) models.NHBuilding {
	building := models.NHBuilding{
		Key:         "Øystein Sunde",
		Name:        "Øystein Sunde",
		ImgUrl:      "oystein_sunde.webp",
		Income:      30,
		Level:       1,
		UpgradeCost: 3000,

		IncomeScale:      1.1,
		UpgradeCostScale: 1.1,
	}

	return *maybeRandomizeStats(&building, stats)
}

func maybeRandomizeStats(building *models.NHBuilding, stats BuildingStats) *models.NHBuilding {
	if stats == RandomStats {
		building.RandomizeStats()
	}
	return building
}
