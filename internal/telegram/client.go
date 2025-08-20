package telegram

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"tea-timer/internal/config"
	"tea-timer/internal/domain/timer"
)

type Client struct {
	Config     *config.Config
	HttpClient *http.Client
	Timers     *timer.Timers
}

func NewClient() *Client {
	client := &Client{
		Config:     config.NewConfig(),
		HttpClient: http.DefaultClient,
		Timers:     timer.NewTimers(),
	}

	client.DeleteWebhook()
	client.SetWebhook()

	return client
}

func (c *Client) SendRequest(method string) int64 {
	queryUrl := c.Config.GetTelegramRequestUrl() + method
	response, err := c.HttpClient.Get(queryUrl)
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

func (c *Client) GetBotInfo() {
	c.SendRequest("getMe")
}

func (c *Client) SetWebhook() {
	c.SendRequest("setWebhook?url=" + c.Config.WebhookUrl)
}

func (c *Client) GetWebhookInfo() {
	c.SendRequest("getWebhookInfo")
}

func (c *Client) DeleteWebhook() {
	c.SendRequest("deleteWebhook?drop_pending_updates")
}

func (c *Client) SendMessage(request TgRequest) int64 {
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

	messageId := c.SendRequest(u)
	return messageId
}

func (c *Client) DeleteMessage(request TgRequest) {
	params := url.Values{}
	params.Set("chat_id", strconv.FormatInt(request.ChatID, 10))
	params.Set("message_id", strconv.FormatInt(request.LastMessageId, 10))

	queryString := params.Encode()

	u := "deleteMessage?" + queryString

	c.SendRequest(u)
}