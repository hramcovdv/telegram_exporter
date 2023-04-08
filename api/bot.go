package api

import (
	"log"

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

	log.Printf("Authorized on account %s", api.Self.UserName)

	return &Bot{api: api}, nil
}

func (b *Bot) Run() error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := b.api.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			if update.Message.Text == "" {
				continue
			}

			var user *types.User

			if user = b.GetUser(update.Message.From.ID); user == nil {
				user = types.NewUser(update.Message.From.ID)
				b.users = append(b.users, user)
			}

			user.Messages++

			// log.Printf("[%d] messages: %d", user.ID, user.Messages)
		}
	}

	return nil
}

func (b *Bot) GetUser(id int64) *types.User {
	for _, user := range b.users {
		if user.ID == id {
			return user
		}
	}

	return nil
}
