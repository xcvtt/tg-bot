package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"tgbot/internal/shitbot"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("shit_bot"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	shitbot.PollUpdates(bot)
}
