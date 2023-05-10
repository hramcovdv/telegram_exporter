package api

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

func NewBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{api: api}, nil
}

func (b *Bot) Run() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := b.api.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

				switch update.Message.Command() {
				case "myid":
					msg.Text = fmt.Sprintf("Your user ID: %d", update.Message.From.ID)
				}

				msg.ReplyToMessageID = update.Message.MessageID
				b.api.Send(msg)
				continue
			}

			userMessages.WithLabelValues(
				fmt.Sprintf("%d", update.Message.From.ID), // label: userid
				fmt.Sprintf("%d", update.Message.Chat.ID), // label: chatid
			).Inc()
		}
	}
}

func (b *Bot) BotName() string {
	return b.api.Self.UserName
}
