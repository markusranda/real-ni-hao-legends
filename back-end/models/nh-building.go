package models

type NHBuilding struct {
	Key         string  `json:"key"`
	Name        string  `json:"name"`
	ImgUrl      string  `json:"imgUrl"`
	Income      float64 `json:"income"`
	Level       float64 `json:"level"`
	BaseCost    float64 `json:"baseCost"`
	UpgradeCost float64 `json:"upgradeCost"`
}
