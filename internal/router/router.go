package router

import (
	"tea-timer/internal/bot"
	"tea-timer/internal/domain/text"
	"tea-timer/internal/telegram"
)

type Route struct {
	Case   string
	Action bot.ActionFunc
}

type Router struct {
	routes []Route
}

func NewRouter(bot *bot.Bot) *Router {
	return &Router{
		routes: routes(bot),
	}
}

func (r *Router) Handle(request telegram.TgRequest) {
	for _, route := range r.routes {
		if request.Text == route.Case {
			route.Action(request)
		}
	}
}

func routes(b *bot.Bot) []Route {
	routes := []Route{
		{Case: "/start", Action: b.StartAction},
		{Case: text.Text("info-button"), Action: b.InfoAction},
		{Case: text.Text("start-button"), Action: b.StartTeaPartyAction},
		{Case: text.Text("start-timer-button"), Action: b.StartTimer},
		{Case: text.Text("reset-timer-button"), Action: b.ResetTimer},
		{Case: text.Text("plus-timer-button"), Action: b.PlusTimer},
		{Case: text.Text("minus-timer-button"), Action: b.MinusTimer},
		{Case: text.Text("first-info-button"), Action: b.FirstInfoAction},
	}

	return routes
}