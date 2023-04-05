package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/serenite11/tg-bot-downloader-from-youtube/pkg/telegram"
	"log"
	"os"
)

func main() {
	error := godotenv.Load()
	if error != nil {
		error.Error()
		return
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
