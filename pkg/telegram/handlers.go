package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	switch message.Command() {
	case "start":
		msg.Text = "Введите url"
		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	path := b.Download(msg.Text)
	msg.Text = b.UrlIsValid(message.Text)
	b.bot.Send(msg)
	if msg.Text == "Your url not from youtube" {
		return
	}
	audioConfig := tgbotapi.NewAudioUpload(message.Chat.ID, path)
	audioConfig.ReplyToMessageID = message.MessageID
	b.bot.Send(audioConfig)
}
