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

			var lvs = []string{
				fmt.Sprintf("%d", update.Message.From.ID), // label: userid
				fmt.Sprintf("%d", update.Message.Chat.ID), // label: chatid
			}

			if update.Message.Animation != nil {
				userAnimations.WithLabelValues(lvs...).Inc()
			}

			if update.Message.Audio != nil {
				userAudios.WithLabelValues(lvs...).Inc()
			}

			if update.Message.Document != nil {
				userDocuments.WithLabelValues(lvs...).Inc()
			}

			if update.Message.Photo != nil {
				userPhotos.WithLabelValues(lvs...).Inc()
			}

			if update.Message.Sticker != nil {
				userStickers.WithLabelValues(lvs...).Inc()
			}

			if update.Message.Video != nil {
				userVideos.WithLabelValues(lvs...).Inc()
			}

			if update.Message.VideoNote != nil {
				userVideoNotes.WithLabelValues(lvs...).Inc()
			}

			if update.Message.Voice != nil {
				userVoices.WithLabelValues(lvs...).Inc()
			}

			userMessages.WithLabelValues(lvs...).Inc()
		}
	}
}

func (b *Bot) BotName() string {
	return b.api.Self.UserName
}
