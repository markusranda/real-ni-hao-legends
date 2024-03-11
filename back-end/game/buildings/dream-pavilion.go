package buildings

import "ni-hao-legends/models"

func CreateDreamPavilion() models.NHBuilding {
	return models.NHBuilding{
		Key:         "Dream pavilion",
		Name:        "Dream pavilion",
		ImgUrl:      "dream_pavilion.webp",
		Income:      100,
		Level:       1,
		BaseCost:    50,
		UpgradeCost: 50,
	}
}
