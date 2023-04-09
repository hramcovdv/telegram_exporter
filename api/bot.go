package api

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/hramcovdv/telegram_exporter/types"
)

type Bot struct {
	api   *tgbotapi.BotAPI
	users []*types.User
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
			if update.Message.Text == "" {
				continue
			}

			userid := update.Message.From.ID

			if update.Message.IsCommand() {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

				switch update.Message.Command() {
				case "myid":
					msg.Text = fmt.Sprintf("Your user ID: %d", userid)
				}

				msg.ReplyToMessageID = update.Message.MessageID

				b.api.Send(msg)

				continue
			}

			var user *types.User

			if user = b.GetUser(userid); user == nil {
				user = types.NewUser(userid)
				b.users = append(b.users, user)
			}

			user.Messages++

			// log.Printf("userid %d, messages: %d", user.ID, user.Messages)
		}
	}
}

func (b *Bot) SelfName() string {
	return b.api.Self.UserName
}

func (b *Bot) GetUser(id int64) *types.User {
	for _, user := range b.users {
		if user.ID == id {
			return user
		}
	}

	return nil
}
