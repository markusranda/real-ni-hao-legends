package game

import "ni-hao-legends/models"

func Chat(command models.Command) {
	message := command.Data["message"].(string)
	msgType := command.Data["type"].(string)
	user := State.Players[command.PlayerId].Town.Name

	newMessage := models.ChatMessage{
		Message:  message,
		UserName: user,
		UserId:   command.PlayerId,
		Type:     msgType,
	}

	if len(State.Chat.Messages) >= 100 {
		// Remove the oldest message
		State.Chat.Messages = State.Chat.Messages[1:]
	}

	// Append the new message
	State.Chat.Messages = append(State.Chat.Messages, newMessage)
}
