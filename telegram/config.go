package telegram

import "os"

type Config struct {
	Token      string
	BaseUrl    string
	WebhookUrl string
}

func NewConfig() *Config {
	return &Config{
		Token:      os.Getenv("TELEGRAM_BOT_TOKEN"),
		BaseUrl:    os.Getenv("TELEGRAM_BOT_URL"),
		WebhookUrl: os.Getenv("TELEGRAM_WEBHOOK_URL"),
	}
}

func (c *Config) getTelegramRequestUrl() string {
	return c.BaseUrl + c.Token + "/"
}
