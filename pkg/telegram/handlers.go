package telegram

import (
	"log"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tank130701/url-shortener-telegram-bot/pkg/helpers"
)



func (b* Bot) handleCommand(message *tgbotapi.Message){
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")
	switch message.Command(){
	case "start":
		msg.Text =  "Этот бот сокращает ссылки"
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}

}

func (b* Bot) handleMessage(message *tgbotapi.Message){
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	// msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	// msg.ReplyToMessageID = message.MessageID
	text := message.Text
    // Проверяем, содержит ли сообщение валидную ссылку
	if containsValidLink(text) {
		replyText := "Сообщение содержит валидную ссылку!"
		msg := tgbotapi.NewMessage(message.Chat.ID, replyText)
		b.bot.Send(msg) 
		shortLink := helpers.GenerateShortLink(text)
		msg = tgbotapi.NewMessage(message.Chat.ID, shortLink)
		b.bot.Send(msg) 
	} else {
		replyText := "Сообщение не содержит валидную ссылку."
		msg := tgbotapi.NewMessage(message.Chat.ID, replyText)
		b.bot.Send(msg)
	}
}

func containsValidLink(text string) bool {
    // Паттерн для проверки ссылки
    linkPattern := regexp.MustCompile(`^(http|https)://[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(\/\S*)?$`)

    // Проверяем, соответствует ли текст сообщения паттерну
    return linkPattern.MatchString(text)
}
