package telegram

import (
	"encoding/json"
	"errors"
	"log"
)

type TgRequest struct {
	ChatID        int64
	LastMessageId int64
	Text          string
	IsCallback    bool
	ParseMode     string
	ReplyMarkup   interface{} `json:"reply_markup,omitempty"`
}

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

func (r *TgRequest) setMarkupButtons(buttons []string) {
	var keyboard [][]KeyboardButton
	for _, buttonText := range buttons {
		keyboard = append(keyboard, []KeyboardButton{{Text: buttonText}})
	}

	r.ReplyMarkup = ReplyKeyboardMarkup{
		Keyboard:       keyboard,
		ResizeKeyboard: true,
	}
}

func (r *TgRequest) removeMarkupButtons() {
	r.ReplyMarkup = ReplyKeyboardRemove{
		RemoveKeyboard: true,
	}
}
