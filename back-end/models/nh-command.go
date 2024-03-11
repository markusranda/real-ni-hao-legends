package models

type Command struct {
	Action   string                 `json:"type"`
	PlayerId PlayerId               `json:"userId"`
	Data     map[string]interface{} `json:"data"`
}
