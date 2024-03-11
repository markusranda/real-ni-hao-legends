package models

type NHMessage struct {
	Message string
	Type    string
	User    string
}

type WSChatEvent struct {
	Message string
	User    string
}

type WSUpdateEvent struct {
	Message string
	User    string
}
