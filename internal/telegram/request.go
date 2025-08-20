package telegram

import (
	"encoding/json"
	"errors"
	"log"
)

func NewTgRequest(body []byte) (TgRequest, error) {
	var webhookMessage WebhookMessage
	if err := json.Unmarshal(body, &webhookMessage); err == nil && webhookMessage.Message.From.Username != "" {

		log.Printf("Received webhook message: %+v\n", webhookMessage)

		return TgRequest{
			LastMessageId: webhookMessage.Message.MessageID,
			ChatID:        webhookMessage.Message.Chat.ID,
			Text:          webhookMessage.Message.Text,
		}, nil
	}

	return TgRequest{}, errors.New("unable to parse request body")
}