package buildings

import (
	"github.com/google/uuid"
	"ni-hao-legends/models"
)

func NewBuilding(key string, name string, imgUrl string, income float64, upgradeCost float64, incomeScale float64, upgradeCostScale float64, stats BuildingStats, rarity models.Rarity) models.NHBuilding {
	building := models.NHBuilding{
		Level:       1,
		UniqueId:    uuid.New(),
		Key:         key,
		Name:        name,
		ImgUrl:      imgUrl,
		Income:      income,
		UpgradeCost: upgradeCost,
		Rarity:      int(rarity),

		IncomeScale:      incomeScale,
		UpgradeCostScale: upgradeCostScale,

		EnhancementSlots:     5,
		UsedEnhancementSlots: 0,
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
		10,   // income
		1000, // upgradeCost
		1.1,  // incomeScale
		1.1,  // upgradeCostScale
		stats,
		models.Common,
	)
}

func CreateHoleInTheGround(stats BuildingStats) models.NHBuilding {
	return NewBuilding(
		"Hole in the ground",
		"Hole in the ground",
		"hole_in_the_ground.webp",
		-20,
		3000,
		1.1,
		.95,
		stats,
		models.Uncommon,
	)
}

func CreateNMSGjenbruk(stats BuildingStats) models.NHBuilding {
	return NewBuilding(
		"NMS Gjenbruk",
		"NMS Gjenbruk",
		"nms_gjenbruk.webp",
		60,   // income
		7200, // upgradeCost
		1.3,  // incomeScale
		1.3,  // upgradeCostScale
		stats,
		models.Rare,
	)
}

func CreateOrientalDragon(stats BuildingStats) models.NHBuilding {
	return NewBuilding(
		"Oriental Dragon",
		"Oriental Dragon",
		"oriental_dragon.webp",
		100,    // income
		400000, // upgradeCost
		1.4,    // incomeScale
		1.4,    // upgradeCostScale
		stats,
		models.Rare,
	)
}

func CreateEricClapton(stats BuildingStats) models.NHBuilding {
	return NewBuilding(
		"Eric Clapton",
		"Eric Clapton",
		"eric_clapton.png",
		15,   // unique income
		1500, // unique upgradeCost
		1.2,  // unique incomeScale
		1.2,  // unique upgradeCostScale
		stats,
		models.Common,
	)
}

func CreateØysteinSunde(stats BuildingStats) models.NHBuilding {
	return NewBuilding(
		"Øystein Sunde",
		"Øystein Sunde",
		"sunde.png",
		150,   // income
		12000, // upgradeCost
		1.5,   // incomeScale
		1.5,   // upgradeCostScale
		stats,
		models.Epic,
	)
}

func CreateKjellElvis(stats BuildingStats) models.NHBuilding {
	return NewBuilding(
		"Kjell Elvis",
		"Kjell Elvis",
		"kjell_elvis.png",
		30,   // income
		3600, // upgradeCost
		1.2,  // incomeScale
		1.2,  // upgradeCostScale
		stats,
		models.Uncommon,
	)
}

func CreateFruitCasket(stats BuildingStats) models.NHBuilding {
	return NewBuilding(
		"Fruit Casket",
		"Fruit Casket",
		"fruit_casket.png",
		200,   // income
		15000, // upgradeCost
		1.6,   // incomeScale
		1.6,   // upgradeCostScale
		stats,
		models.Legendary,
	)
}

func maybeRandomizeStats(building *models.NHBuilding, stats BuildingStats) *models.NHBuilding {
	if stats == RandomStats {
		building.RandomizeStats()
	}
	return building
}
