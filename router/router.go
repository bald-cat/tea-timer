package router

import (
	"sync"
	"tgbot/telegram"
	"tgbot/texts"
)

type Route struct {
	Case   string
	Action telegram.ActionFunc
}

type Router struct {
	routes []Route
}

var (
	once     sync.Once
	instance *Router
)

func NewRouter(telegram *telegram.Telegram, request telegram.TgRequest) {
	once.Do(func() {
		instance = &Router{
			routes: routes(telegram),
		}
	})

	instance.Handle(request)
}

func (r *Router) Handle(request telegram.TgRequest) {
	for _, route := range r.routes {
		if request.Text == route.Case {
			route.Action(request)
		}
	}
}

func routes(telegram *telegram.Telegram) []Route {
	routes := []Route{
		{Case: "/start", Action: telegram.StartAction},
		{Case: texts.Text("info-button"), Action: telegram.InfoAction},
		{Case: texts.Text("start-button"), Action: telegram.StartTeaPartyAction},
		{Case: texts.Text("start-timer-button"), Action: telegram.StartTimer},
		{Case: texts.Text("reset-timer-button"), Action: telegram.ResetTimer},
		{Case: texts.Text("plus-timer-button"), Action: telegram.PlusTimer},
		{Case: texts.Text("minus-timer-button"), Action: telegram.MinusTimer},
		{Case: texts.Text("first-info-button"), Action: telegram.FirstInfoAction},
	}

	return routes
}
