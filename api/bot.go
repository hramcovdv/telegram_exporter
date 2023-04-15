package api

import (
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

			if update.Message.From.IsBot {
				continue
			}

			var (
				userid = update.Message.From.ID
				chatid = update.Message.Chat.ID
			)

			var user *types.User
			if user = b.GetUser(userid, chatid); user == nil {
				user = types.NewUser(userid, chatid)
				b.users = append(b.users, user)
			}

			if update.Message.IsCommand() {
				msg := tgbotapi.NewMessage(chatid, "")

				switch update.Message.Command() {
				case "myid":
					msg.Text = user.GetMyIdMsg()
				case "mystats":
					msg.Text = user.GetMyStatsMsg()
				}

				msg.ReplyToMessageID = update.Message.MessageID

				b.api.Send(msg)

				continue
			}

			user.IncMessages(update.Message.Text)
		}
	}
}

func (b *Bot) SelfName() string {
	return b.api.Self.UserName
}

func (b *Bot) GetUser(id int64, chatid int64) *types.User {
	for _, user := range b.users {
		if user.ID == id && user.ChatID == chatid {
			return user
		}
	}

	return nil
}

func (b *Bot) GetUsers() []*types.User {
	return b.users
}
