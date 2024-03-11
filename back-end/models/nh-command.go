package models

type Command struct {
	Action   string                 `json:"type"`
	PlayerId PlayerId               `json:"userId"`
	Data     map[string]interface{} `json:"data"`
}

type BuildingCommand struct {
	*Command
	Key    string `json:"key"`
	Action string `json:"action"`
}
