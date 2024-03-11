package models

type GameState struct {
	Players  map[PlayerId]*PlayerGameState
	Chat     Chat
	Samba    NHSamba
	Commands []Command
}
type PlayerGameState struct {
	RetirementFund float64   `json:"retirementFund"`
	Inventory      Inventory `json:"inventory"`
	Town           NHTown    `json:"town"`
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
