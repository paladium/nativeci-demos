package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func getRandomPhoto() ([]byte, error) {
	r, err := http.Get("https://source.unsplash.com/random")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case "inspiration":
			photo, err := getRandomPhoto()
			if err != nil {
				panic(err)
			}
			photoMessage := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, tgbotapi.FileBytes{
				Bytes: photo,
			})
			_, err = bot.Send(photoMessage)
			if err != nil {
				panic(err)
			}
		default:
			msg.Text = "No handler"
			bot.Send(msg)
		}
	}
}
