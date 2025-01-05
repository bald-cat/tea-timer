package telegram

import "tgbot/texts"

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type ReplyKeyboardMarkup struct {
	Keyboard       [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard bool               `json:"resize_keyboard"`
}

type KeyboardButton struct {
	Text string `json:"text"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
}

type ForceReply struct {
	ForceReply bool `json:"force_reply"`
}

func GetStartButtons() []string {
	return []string{
		texts.Text("start-button"),
		texts.Text("info-button"),
		texts.Text("first-info-button"),
	}
}

func GetBaseButtons() []string {
	return []string{
		texts.Text("start-timer-button"),
		texts.Text("plus-timer-button"),
		texts.Text("minus-timer-button"),
		texts.Text("reset-timer-button"),
	}
}
