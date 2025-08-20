package bot

import (
	"fmt"
	"log"
	"strconv"
	"tea-timer/internal/domain/text"
	"tea-timer/internal/telegram"
	"time"
)

func (b *Bot) StartAction(request telegram.TgRequest) {
	request.Text = text.Text("start")

	buttons := telegram.GetStartButtons()

	request.SetMarkupButtons(buttons)
	b.telegramClient.SendMessage(request)
}

func (b *Bot) InfoAction(request telegram.TgRequest) {
	request.Text = text.Text("info")

	buttons := telegram.GetStartButtons()

	request.SetMarkupButtons(buttons)
	b.telegramClient.SendMessage(request)
}

func (b *Bot) FirstInfoAction(request telegram.TgRequest) {
	request.Text = text.Text("first-info")

	buttons := telegram.GetStartButtons()

	request.SetMarkupButtons(buttons)
	b.telegramClient.SendMessage(request)
}

func (b *Bot) StartTeaPartyAction(request telegram.TgRequest) {
	request.Text = text.Text("start-timer")

	buttons := telegram.GetBaseButtons()

	request.SetMarkupButtons(buttons)
	b.telegramClient.SendMessage(request)
}

func (b *Bot) StartTimer(request telegram.TgRequest) {
	timerByChatId := b.telegramClient.Timers.GetTimerByChatId(request.ChatID)
	request.Text = fmt.Sprintf("Таймер установлен на %v", timerByChatId.Duration)
	request.RemoveMarkupButtons()
	request.LastMessageId = b.telegramClient.SendMessage(request)
	duration := timerByChatId.Duration
	timerByChatId.AddDuration(time.Second * 15)
	time.AfterFunc(duration, func() {
		request.Text = fmt.Sprintf(text.Text("timer-end"), timerByChatId.Duration)

		buttons := telegram.GetBaseButtons()

		request.SetMarkupButtons(buttons)
		log.Printf(strconv.FormatInt(request.LastMessageId, 10))
		b.telegramClient.DeleteMessage(request)
		b.telegramClient.SendMessage(request)
	})
}

func (b *Bot) ResetTimer(request telegram.TgRequest) {
	timerByChatId := b.telegramClient.Timers.GetTimerByChatId(request.ChatID)
	timerByChatId.ResetDuration()
	request.Text = text.Text("reset")

	buttons := telegram.GetStartButtons()

	request.SetMarkupButtons(buttons)
	b.telegramClient.SendMessage(request)
}

func (b *Bot) PlusTimer(request telegram.TgRequest) {
	timerByChatId := b.telegramClient.Timers.GetTimerByChatId(request.ChatID)
	timerByChatId.AddDuration(time.Second * 5)
	request.Text = fmt.Sprintf("Время увеличено, следующий таймер будет запущен на  %v", timerByChatId.Duration)

	buttons := telegram.GetBaseButtons()

	request.SetMarkupButtons(buttons)
	b.telegramClient.SendMessage(request)
}

func (b *Bot) MinusTimer(request telegram.TgRequest) {
	timerByChatId := b.telegramClient.Timers.GetTimerByChatId(request.ChatID)
	timerByChatId.MinusDuration(time.Second * 5)
	request.Text = fmt.Sprintf("Время уменьшено, следующий таймер будет запущен на  %v", timerByChatId.Duration)

	buttons := telegram.GetBaseButtons()

	request.SetMarkupButtons(buttons)
	b.telegramClient.SendMessage(request)
}