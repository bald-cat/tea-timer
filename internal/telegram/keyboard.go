package telegram

import "tea-timer/internal/domain/text"

func GetStartButtons() []string {
	return []string{
		text.Text("start-button"),
		text.Text("info-button"),
		text.Text("first-info-button"),
	}
}

func GetBaseButtons() []string {
	return []string{
		text.Text("start-timer-button"),
		text.Text("plus-timer-button"),
		text.Text("minus-timer-button"),
		text.Text("reset-timer-button"),
	}
}

func (r *TgRequest) SetMarkupButtons(buttons []string) {
	var keyboard [][]KeyboardButton
	for _, buttonText := range buttons {
		keyboard = append(keyboard, []KeyboardButton{{Text: buttonText}})
	}

	r.ReplyMarkup = ReplyKeyboardMarkup{
		Keyboard:       keyboard,
		ResizeKeyboard: true,
	}
}

func (r *TgRequest) RemoveMarkupButtons() {
	r.ReplyMarkup = ReplyKeyboardRemove{
		RemoveKeyboard: true,
	}
}