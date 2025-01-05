package timer

import (
	"log"
	"time"
)

const StartDuration = time.Second * 15

type Timers struct {
	Timers map[int64]*Timer
}

func NewTimers() *Timers {
	return &Timers{
		Timers: make(map[int64]*Timer),
	}
}

func (t *Timers) GetTimerByChatId(chatId int64) *Timer {
	if timer, exists := t.Timers[chatId]; exists {
		log.Printf("%+v", timer)
		return timer
	}

	newTimer := &Timer{
		ChatId:   chatId,
		Period:   1,
		Duration: StartDuration,
	}

	t.Timers[chatId] = newTimer

	return newTimer
}

func (t *Timers) AddDurationByChatId(chatId int64, duration time.Duration) {
	timer := t.Timers[chatId]
	timer.Duration += duration
	t.Timers[chatId] = timer
}

type Timer struct {
	ChatId   int64
	Period   int
	Duration time.Duration
}

func (t *Timer) IncrementPeriod() {
	t.Period++
}

func (t *Timer) DecrementPeriod() {
	t.Period--
}

func (t *Timer) AddDuration(duration time.Duration) {
	t.Duration += duration
}

func (t *Timer) MinusDuration(duration time.Duration) {
	t.Duration -= duration
}

func (t *Timer) ResetDuration() {
	t.Duration = StartDuration
}
