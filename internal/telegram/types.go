package telegram

type WebhookMessage struct {
	Message struct {
		MessageID int64 `json:"message_id"`
		From      struct {
			Username string `json:"username"`
		} `json:"from"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
		Text string `json:"text"`
	} `json:"message"`
}

type Response struct {
	OK     bool `json:"ok"`
	Result struct {
		MessageID int64 `json:"message_id"`
	}
}

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

type TgRequest struct {
	ChatID        int64
	LastMessageId int64
	Text          string
	IsCallback    bool
	ParseMode     string
	ReplyMarkup   interface{} `json:"reply_markup,omitempty"`
}