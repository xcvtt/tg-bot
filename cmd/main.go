package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"os"
	"sort"
)

type user struct {
	id   int64
	name string
	hp   int
}

func getUserById(users *[]user, userId int64) (user, bool) {
	for _, u := range *users {
		if u.id == userId {
			return u, true
		}
	}

	return user{}, false
}

func tryShit(users *[]user, userInfo *tgbotapi.User) string {
	usr, ok := getUserById(users, userInfo.ID)

	if !ok {
		usr = user{
			id:   userInfo.ID,
			name: userInfo.FirstName,
			hp:   100,
		}

		*users = append(*users, usr)
	}

	if usr.hp <= 0 {
		return "Ты не можешь срать на головы, твое hp равно 0"
	}

	r := rand.Intn(len(*users))

	var msg string

	if (*users)[r].id == usr.id {
		msg = fmt.Sprintf("%s насрал себе на голову", usr.name)
	} else {
		msg = fmt.Sprintf("%s насрал на голову %s", usr.name, (*users)[r].name)
	}

	dmg := rand.Intn(5) + 1
	(*users)[r].hp -= dmg

	hpMsg := fmt.Sprintf("Урон %d. Осталось hp: %d", dmg, (*users)[r].hp)

	return fmt.Sprintf("%s\n%s", msg, hpMsg)
}

func getHp(users []user) string {
	if len(users) == 0 {
		return "Users empty"
	}

	sort.SliceStable(users, func(i, j int) bool {
		return users[i].hp > users[j].hp
	})

	var msg string

	for i, u := range users {
		msg += fmt.Sprintf("%d. %s осталось hp: %d \n", i+1, u.name, u.hp)
	}

	return msg
}

const M07 = -1001552222270

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("shit_bot"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	users := make([]user, 0, 30)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if update.Message.Chat.ID != M07 {
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "nasrat":
			msg.Text += tryShit(&users, update.Message.From)
		case "hp":
			msg.Text += getHp(users)
		default:
		}

		if _, err := bot.Send(msg); err != nil {
			log.Println("Gavneco chet yopta")
		}
	}
}
