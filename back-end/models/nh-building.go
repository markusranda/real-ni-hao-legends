package models

import (
	"fmt"
	"github.com/google/uuid"
	"math"
	"math/rand"
)

type NHBuilding struct {
	Key         string    `json:"key"`
	UniqueId    uuid.UUID `json:"uniqueId"`
	Name        string    `json:"name"`
	ImgUrl      string    `json:"imgUrl"`
	Rarity      int       `json:"rarity"`
	Income      float64   `json:"income"`
	Level       float64   `json:"level"`
	UpgradeCost float64   `json:"upgradeCost"`

	// scaling parameters
	IncomeScale      float64 `json:"incomeScale"`
	BaseCostScale    float64 `json:"baseCostScale"`
	UpgradeCostScale float64 `json:"upgradeCostScale"`

	// enhancement
	EnhancementSlots     int `json:"enhancementSlots"`
	UsedEnhancementSlots int `json:"usedEnhancementSlots"`
}

type Rarity int

const (
	Common Rarity = iota
	Uncommon
	Rare
	Epic
	Legendary
)

func (n *NHBuilding) RandomizeStats() {
	n.Income = math.Floor(n.Income - n.Income*(rand.Float64()-0.5))
	n.UpgradeCost = math.Floor(n.UpgradeCost - n.UpgradeCost*(rand.Float64()-0.5))

	// todo create working randomizer when i am well rested
	//n.IncomeScale = math.Floor(n.IncomeScale-n.IncomeScale*((rand.Float64()/-0.5)/5)*100) / 100
	//n.UpgradeCostScale = math.Floor(n.UpgradeCostScale-n.UpgradeCostScale*((rand.Float64()/-0.5)/5)*100) / 100
}

func (n *NHBuilding) Upgrade() {
	n.Income = math.Floor(n.Income * n.IncomeScale)
	n.UpgradeCost = math.Floor(n.UpgradeCost * n.UpgradeCostScale)

	n.Level++
}

type Buildings []NHBuilding

func (b Buildings) FindById(key string) (*NHBuilding, error) {
	for _, building := range b {
		if building.Key == key {
			return &building, nil
		}
	}
	return nil, fmt.Errorf("building with id %s not found", key)
}
