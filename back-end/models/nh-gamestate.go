package models

import "github.com/google/uuid"

type GameState struct {
	Players  map[PlayerId]*PlayerGameState
	Chat     Chat
	Samba    NHSamba
	Commands []Command
}
type PlayerGameState struct {
	RetirementFund float64    `json:"retirementFund"`
	Effects        []NHEffect `json:"effects"`
	Inventory      Inventory  `json:"inventory"`
	Town           NHTown     `json:"town"`
}

type Inventory struct {
	Buildings []NHBuilding `json:"buildings"`
}

type Chat struct {
	Messages []ChatMessage
}

type ChatMessage struct {
	Message  string `json:"message"`
	UserName string `json:"userName"`
	UserId   string `json:"userId"`
	Type     string `json:"type"`
}

type NHEffect struct {
	EffectType string    `json:"effectType"`
	Timestamp  int64     `json:"timestamp"`
	UUID       uuid.UUID `json:"uuid"`
}
