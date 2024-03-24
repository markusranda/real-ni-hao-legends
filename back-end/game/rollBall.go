package game

import (
	"github.com/google/uuid"
	"ni-hao-legends/models"
)

func RollBall(command models.Command) error {
	for _, player := range State.Players {
		player.Effects = append(player.Effects, models.NHEffect{
			EffectType: "rollBall",
			Timestamp:  command.Time,
			UUID:       uuid.New(),
		})
	}
	return nil
}
