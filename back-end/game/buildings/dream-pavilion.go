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
