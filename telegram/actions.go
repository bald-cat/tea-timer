package telegram

import (
	"fmt"
	"log"
	"strconv"
	"tgbot/texts"
	"time"
)

type ActionFunc func(request TgRequest)

func (t *Telegram) StartAction(request TgRequest) {
	request.Text = texts.Text("start")

	buttons := GetStartButtons()

	request.setMarkupButtons(buttons)
	t.SendMessage(request)
}

func (t *Telegram) InfoAction(request TgRequest) {
	request.Text = texts.Text("info")
	t.SendMessage(request)
}

func (t *Telegram) StartTeaPartyAction(request TgRequest) {
	request.Text = texts.Text("start-timer")
	buttons := GetBaseButtons()
	request.setMarkupButtons(buttons)
	t.SendMessage(request)
}

func (t *Telegram) StartTimer(request TgRequest) {
	timerByChatId := t.Timers.GetTimerByChatId(request.ChatID)
	request.Text = fmt.Sprintf("Таймер установлен на %v", timerByChatId.Duration)
	request.removeMarkupButtons()
	request.LastMessageId = t.SendMessage(request)
	duration := timerByChatId.Duration
	timerByChatId.AddDuration(time.Second * 15)
	time.AfterFunc(duration, func() {
		request.Text = fmt.Sprintf("Время вышло, следующий таймер будет запущен на  %v", timerByChatId.Duration)

		buttons := GetBaseButtons()

		request.setMarkupButtons(buttons)
		log.Printf(strconv.FormatInt(request.LastMessageId, 10))
		t.DeleteMessage(request)
		t.SendMessage(request)
	})
}

func (t *Telegram) ResetTimer(request TgRequest) {
	timerByChatId := t.Timers.GetTimerByChatId(request.ChatID)
	timerByChatId.ResetDuration()
	request.Text = fmt.Sprintf("Таймер сброшен, следующий таймер будет запущен на  %v", timerByChatId.Duration)

	buttons := GetBaseButtons()

	request.setMarkupButtons(buttons)
	t.SendMessage(request)
}

func (t *Telegram) PlusTimer(request TgRequest) {
	timerByChatId := t.Timers.GetTimerByChatId(request.ChatID)
	timerByChatId.AddDuration(time.Second * 5)
	request.Text = fmt.Sprintf("Время увеличено, следующий таймер будет запущен на  %v", timerByChatId.Duration)

	buttons := GetBaseButtons()

	request.setMarkupButtons(buttons)
	t.SendMessage(request)
}

func (t *Telegram) MinusTimer(request TgRequest) {
	timerByChatId := t.Timers.GetTimerByChatId(request.ChatID)
	timerByChatId.MinusDuration(time.Second * 5)
	request.Text = fmt.Sprintf("Время уменьшено, следующий таймер будет запущен на  %v", timerByChatId.Duration)

	buttons := GetBaseButtons()

	request.setMarkupButtons(buttons)
	t.SendMessage(request)
}
