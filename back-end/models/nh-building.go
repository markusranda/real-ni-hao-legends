package models

import (
	"math/rand"
)

type NHBuilding struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	ImgUrl string `json:"imgUrl"`

	Income      float64 `json:"income"`
	Level       float64 `json:"level"`
	UpgradeCost float64 `json:"upgradeCost"`

	// scaling parameters
	IncomeScale      float64 `json:"incomeScale"`
	BaseCostScale    float64 `json:"baseCostScale"`
	UpgradeCostScale float64 `json:"upgradeCostScale"`

	// enhancement
	EnhancementSlots     int `json:"enhancementSlots"`
	UsedEnhancementSlots int `json:"usedEnhancementSlots"`
}

func (n *NHBuilding) RandomizeStats() {
	n.Income = n.Income - n.Income*(rand.Float64()-0.5)
	n.UpgradeCost = n.UpgradeCost - n.UpgradeCost*(rand.Float64()-0.5)

	n.IncomeScale = n.IncomeScale - n.IncomeScale*((rand.Float64()/-0.5)/5)
	n.UpgradeCostScale = n.UpgradeCostScale - n.UpgradeCostScale*((rand.Float64()/-0.5)/5)
}

func (n *NHBuilding) Upgrade() {
	n.Income = n.Income * n.IncomeScale
	n.UpgradeCost = n.UpgradeCost * n.UpgradeCostScale

	n.Level++
}
