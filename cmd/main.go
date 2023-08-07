package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/tank130701/url-shortener-telegram-bot/pkg/telegram"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	BotToken := os.Getenv("BOT_TOKEN")
	botAPI, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Panic(err)
	}

	// botAPI.Debug = true


	telegramBot := telegram.NewBot(botAPI)
	if err := telegramBot.Run(); err != nil{
		log.Fatal(err)
	}
	log.Print("TelegramBot Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("TelegramBot Shutting Down")

}
