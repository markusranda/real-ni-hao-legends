package models

type Command struct {
	Action   string                 `json:"type"`
	PlayerId PlayerId               `json:"userId"`
	Time     int64                  `json:"time"`
	Data     map[string]interface{} `json:"data"`
}
