package buildings

import (
	"math/rand"
	"ni-hao-legends/models"
)

func CreateDreamPavilion(
	stats int,
) models.NHBuilding {
	if stats == RANDOM {
		return models.NHBuilding{
			Key:         "Dream pavilion",
			Name:        "Dream pavilion",
			ImgUrl:      "dream_pavilion.webp",
			Income:      float64(rand.Intn(10) + 1),
			Level:       1,
			BaseCost:    float64(rand.Intn(100) + 50),
			UpgradeCost: float64(rand.Intn(1000) + 500),
		}
	} else if stats == INIT {
		return models.NHBuilding{
			Key:         "Dream pavilion",
			Name:        "Dream pavilion",
			ImgUrl:      "dream_pavilion.webp",
			Income:      5,
			Level:       1,
			BaseCost:    50,
			UpgradeCost: 500,
		}
	} else {
		panic("get it together, you can do better than this")
	}
}

func CreateHoleInTheGround(stats int) models.NHBuilding {
	if stats == RANDOM {
		return models.NHBuilding{
			Key:         "Hole in the ground",
			Name:        "Hole in the ground",
			ImgUrl:      "hole_in_the_ground.webp",
			Income:      float64(rand.Intn(20) + 1), // Random number between 1 and 20
			Level:       1,
			BaseCost:    float64(rand.Intn(200) + 100),   // Random number between 100 and 300
			UpgradeCost: float64(rand.Intn(2000) + 1000), // Random number between 1000 and 3000
		}
	} else if stats == INIT {
		return models.NHBuilding{
			Key:         "Hole in the ground",
			Name:        "Hole in the ground",
			ImgUrl:      "hole_in_the_ground.webp",
			Income:      10,
			Level:       1,
			BaseCost:    100,
			UpgradeCost: 1000,
		}
	} else {
		panic("get it together, you can do better than this")
	}
}

func CreateNMSGjenbruk(stats int) models.NHBuilding {
	if stats == RANDOM {
		return models.NHBuilding{
			Key:         "NMS Gjenbruk",
			Name:        "NMS Gjenbruk",
			ImgUrl:      "nms_gjenbruk.webp",
			Income:      float64(rand.Intn(30) + 1), // Random number between 1 and 30
			Level:       1,
			BaseCost:    float64(rand.Intn(300) + 150),   // Random number between 150 and 450
			UpgradeCost: float64(rand.Intn(3000) + 1500), // Random number between 1500 and 4500
		}
	} else if stats == INIT {
		return models.NHBuilding{
			Key:         "NMS Gjenbruk",
			Name:        "NMS Gjenbruk",
			ImgUrl:      "nms_gjenbruk.webp",
			Income:      15,
			Level:       1,
			BaseCost:    150,
			UpgradeCost: 1500,
		}
	} else {
		panic("get it together, you can do better than this")
	}
}

func CreateOrientalDragon(stats int) models.NHBuilding {
	if stats == RANDOM {
		return models.NHBuilding{
			Key:         "Oriental Dragon",
			Name:        "Oriental Dragon",
			ImgUrl:      "oriental_dragon.webp",
			Income:      float64(rand.Intn(40) + 1), // Random number between 1 and 40
			Level:       1,
			BaseCost:    float64(rand.Intn(400) + 200),   // Random number between 200 and 600
			UpgradeCost: float64(rand.Intn(4000) + 2000), // Random number between 2000 and 6000
		}
	} else if stats == INIT {
		return models.NHBuilding{
			Key:         "Oriental Dragon",
			Name:        "Oriental Dragon",
			ImgUrl:      "oriental_dragon.webp",
			Income:      20,
			Level:       1,
			BaseCost:    200,
			UpgradeCost: 2000,
		}
	} else {
		panic("get it together, you can do better than this")
	}
}
