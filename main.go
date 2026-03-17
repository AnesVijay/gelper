package main

import (
	"fmt"
	"gelper/config"
	telegrambot "gelper/telegram-bot"
	"log"

	"github.com/AnesVijay/glogger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	conf := config.New("config.yaml")
	glogger.InitLogger("log", conf.LogLevel)

	tgbot, err := telegrambot.BotInit()
	if err != nil {
		glogger.GetLogger().SendError(fmt.Sprintf("failed to initialize Telegram bot: %v", err))
		log.Fatal(err)
	} else {
		glogger.GetLogger().SendInfo("Telegram bot initialized!")
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := tgbot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start":
			msg.Text = fmt.Sprintf("ChatID: %d", update.Message.Chat.ID)
		}

		if _, err := tgbot.Send(msg); err != nil {
			log.Panic(err)
		}

	}
}
