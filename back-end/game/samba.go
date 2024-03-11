package game

import (
	"ni-hao-legends/models"
)

func ScoreUpdate(command models.Command) {
	userId, ok := command.Data["userId"].(string)
	if !ok {
		panic("Everybody leave the building, there is no userId!!!")
	}
	scoreValue, ok := command.Data["score"].(float64)
	if !ok {
		panic("Everybody leave the building, there is no score!!!")
	}
	townName, ok := command.Data["townName"].(string)
	if !ok {
		panic("Everybody leave the building, there is no townName!!!")
	}

	score := models.NHScore{UserId: userId, Score: scoreValue, TownName: townName}
	State.Samba.Scores = append(State.Samba.Scores, score)
}
