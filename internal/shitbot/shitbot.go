package shitbot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"sort"
	"tgbot/internal/users"
	"tgbot/models"
)

func getUserById(r users.Repository, userId int64) (*models.User, bool) {
	u, err := r.GetById(userId)

	if err != nil {
		return u, true
	}

	return nil, false
}

func tryShit(r users.Repository, userInfo *tgbotapi.User) string {
	usr, ok := getUserById(r, userInfo.ID)

	if !ok {
		usr = &models.User{
			Id:   userInfo.ID,
			Name: userInfo.FirstName,
			Hp:   100,
		}

		err := r.Create(usr)
		if err != nil {
			log.Fatal(err)
		}
	}

	if usr.Hp <= 0 {
		return "Ты не можешь срать на головы, твое hp равно 0"
	}

	userList, err := r.GetAll()

	if err != nil {
		log.Fatal(err)
	}

	i := rand.Intn(len(userList))

	var msg string

	if userList[i].Id == usr.Id {
		msg = fmt.Sprintf("%s насрал себе на голову", usr.Name)
	} else {
		msg = fmt.Sprintf("%s насрал на голову %s", usr.Name, userList[i].Name)
	}

	dmg := rand.Intn(5) + 1
	userList[i].Hp -= dmg

	hpMsg := fmt.Sprintf("Урон %d. Осталось hp: %d", dmg, userList[i].Hp)

	return fmt.Sprintf("%s\n%s", msg, hpMsg)
}

func getHp(r users.Repository) string {
	usrs, err := r.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	if len(usrs) == 0 {
		return "Users empty"
	}

	sort.SliceStable(usrs, func(i, j int) bool {
		return usrs[i].Hp > usrs[j].Hp
	})

	var msg string

	for i, u := range usrs {
		msg += fmt.Sprintf("%d. %s осталось hp: %d \n", i+1, u.Name, u.Hp)
	}

	return msg
}
