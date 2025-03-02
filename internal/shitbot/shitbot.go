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

func getUserById(r users.Repository, userId int64) *models.User {
	u, err := r.GetById(userId)

	if err != nil {
		return nil
	}

	return u
}

func tryShit(r users.Repository, userInfo *tgbotapi.User, ch chan string) {
	usr := getUserById(r, userInfo.ID)

	if usr == nil {
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
		ch <- "Ты не можешь срать на головы, твое hp равно 0"
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

	dmg := rand.Intn(20) + 1
	userList[i].Hp -= dmg

	err = r.Update(&userList[i])
	if err != nil {
		log.Fatal(err)
	}

	hpMsg := fmt.Sprintf("Урон %d. Осталось hp: %d", dmg, userList[i].Hp)

	ch <- fmt.Sprintf("%s\n%s", msg, hpMsg)
}

func rollHp(r users.Repository, userInfo *tgbotapi.User, ch chan string) {
	usr := getUserById(r, userInfo.ID)

	if usr == nil {
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

	var msg string
	i := rand.Intn(2)
	hp := rand.Intn(200)

	if i%2 == 1 {
		msg = fmt.Sprintf("Гавно тебе на рыло. Следующая попытка через час")
	} else {
		usr.Hp += hp
		msg = fmt.Sprintf("Поздравляю. Ты получил %d hp. Твое текущее hp: %d", hp, usr.Hp)
	}

	var err = r.Update(usr)
	if err != nil {
		log.Fatal(err)
	}

	ch <- msg
}

func getHp(r users.Repository, ch chan string) {
	usrs, err := r.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	if len(usrs) == 0 {
		ch <- "Users empty"
	}

	sort.SliceStable(usrs, func(i, j int) bool {
		return usrs[i].Hp > usrs[j].Hp
	})

	var msg string

	for i, u := range usrs {
		msg += fmt.Sprintf("%d. %s осталось hp: %d \n", i+1, u.Name, u.Hp)
	}

	ch <- msg
}
