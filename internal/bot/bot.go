package bot

import (
	"tea-timer/internal/telegram"
)

type ActionFunc func(request telegram.TgRequest)

type Bot struct {
	telegramClient *telegram.Client
}

func NewBot(client *telegram.Client) *Bot {
	return &Bot{
		telegramClient: client,
	}
}