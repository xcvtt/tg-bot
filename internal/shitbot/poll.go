package shitbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"tgbot/internal/store"
	"tgbot/internal/users"
)

const M07 = -1001552222270

func PollUpdates(bot *tgbotapi.BotAPI) {
	st := store.New()
	err := st.Open()
	if err != nil {
		log.Fatal(err)
	}
	repo := users.NewRepository(st)

	cfg := tgbotapi.NewUpdate(0)
	cfg.Timeout = 30

	updates := bot.GetUpdatesChan(cfg)

	for update := range updates {
		if update.Message == nil || update.Message.Chat.ID != M07 || !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ParseMode = "HTML"

		ch := make(chan string)

		switch update.Message.Command() {
		case "nasrat":
			go tryShit(repo, update.Message.From, ch)
		case "hp":
			go getHp(repo, ch)
		case "azino":
			go rollHp(repo, update.Message.From, ch)
		case "obossat":
			go urinate(repo, update.Message.From, ch)
		default:
			continue
		}

		go func() {
			msg.Text = <-ch
			if _, err := bot.Send(msg); err != nil {
				log.Println("Gavneco chet yopta")
			}
		}()
	}
}
