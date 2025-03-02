package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"tgbot/internal/shitbot"
)

func main() {
	fmt.Println("Starting shitbot...")

	bot, err := tgbotapi.NewBotAPI(os.Getenv("SHITBOT_API"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	shitbot.PollUpdates(bot)
}
