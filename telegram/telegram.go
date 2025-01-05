package telegram

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"tgbot/timer"
)

type Telegram struct {
	Config *Config
	Client *http.Client
	Timers *timer.Timers
}

func NewTelegram() *Telegram {
	tg := &Telegram{
		Config: NewConfig(),
		Client: http.DefaultClient,
		Timers: timer.NewTimers(),
	}

	tg.DeleteWebhook()
	tg.SetWebhook()

	return tg
}

func (t *Telegram) SendRequest(method string) int64 {
	queryUrl := t.getTelegramRequestUrl() + method
	response, err := t.Client.Get(queryUrl)
	if err != nil {
		log.Println("Error getting bot info:", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	log.Println("Body:", string(body))

	if err != nil {
		log.Println("Error reading response body:", err)
	}

	var res Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println("Error unmarshalling response JSON:", err)
	}

	return res.Result.MessageID
}

func (t *Telegram) getTelegramRequestUrl() string {
	return t.Config.BaseUrl + t.Config.Token + "/"
}

func (t *Telegram) GetBotInfo() {
	t.SendRequest("getMe")
}

func (t *Telegram) SetWebhook() {
	t.SendRequest("setWebhook?url=" + t.Config.WebhookUrl)
}

func (t *Telegram) GetWebhookInfo() {
	t.SendRequest("getWebhookInfo")
}

func (t *Telegram) DeleteWebhook() {
	t.SendRequest("deleteWebhook?drop_pending_updates")
}

func (t *Telegram) SendMessage(request TgRequest) int64 {
	params := url.Values{}
	params.Set("chat_id", strconv.FormatInt(request.ChatID, 10))
	params.Set("text", request.Text)

	if request.ReplyMarkup != nil {
		replyMarkupJSON, err := json.Marshal(request.ReplyMarkup)
		if err != nil {
			log.Println("Error serializing reply_markup:", err)
		}
		params.Set("reply_markup", string(replyMarkupJSON))
	}

	queryString := params.Encode()

	log.Println("Sending request:", queryString)

	u := "sendMessage?" + queryString

	messageId := t.SendRequest(u)
	return messageId
}

func (t *Telegram) DeleteMessage(request TgRequest) {
	params := url.Values{}
	params.Set("chat_id", strconv.FormatInt(request.ChatID, 10))
	params.Set("message_id", strconv.FormatInt(request.LastMessageId, 10))

	queryString := params.Encode()

	u := "deleteMessage?" + queryString

	t.SendRequest(u)
}
