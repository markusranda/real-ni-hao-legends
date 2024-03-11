package models

type NHScore struct {
	UserId   string  `json:"userId"`
	Score    float64 `json:"score"`
	TownName string  `json:"townName"`
}

type NHSamba struct {
	Scores []NHScore `json:"scores"`
}
