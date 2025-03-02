package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgbot/internal/shitbot"
)

func main() {
	//bot, err := tgbotapi.NewBotAPI(os.Getenv("shit_bot"))
	bot, err := tgbotapi.NewBotAPI("6548636913:AAFKC2JB8PQzya7VU8QQBUe_hXKvRlGrHXQ")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	shitbot.PollUpdates(bot)
}
