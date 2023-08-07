package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot* tgbotapi.BotAPI) *Bot{
	return&Bot{bot: bot}
}

func (b* Bot) Run()error{
	
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)
	b.handleUpdates(updates)
	
	return nil
}

func (b* Bot) handleUpdates(updates tgbotapi.UpdatesChannel){
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}
		if update.Message.IsCommand() { // ignore any non-command Messages
			b.handleCommand(update.Message)
			continue
		}
		if update.Message != nil { // If we got a message
			b.handleMessage(update.Message)
		}
	}
}

