package telegrambot

import (
	"gelper/config"
	"net/http"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BotInit() (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(config.Conf.Telegram.Token)
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	if len(config.Conf.Telegram.ProxyURL) > 0 {
		proxyURL, _ := url.Parse(config.Conf.Telegram.ProxyURL)
		bot.Client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		}
	}

	return bot, nil
}
