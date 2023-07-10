package commands

import (
	"context"
	"os"
	"os/signal"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/cobra"
	"github.com/wisnuragaprawida/project/bootstrap"
	"github.com/wisnuragaprawida/project/pkg/log"
)

func init() {
	registerCommand(startBot2)
}

func startBot2(dep *bootstrap.Dependency) *cobra.Command {
	return &cobra.Command{
		Use:   "bot2",
		Short: "bot2 service",
		Run: func(cmd *cobra.Command, args []string) {
			dep.GetConfig()

			ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

			bot, err := tgbotapi.NewBotAPI("5126877034:AAFkUGUuS6d-z6TDLe6NIWgLcUUGi73U3to")
			if err != nil {
				log.Panic(err)
			}

			bot.Debug = true

			log.Infof("Authorized on account %s", bot.Self.UserName)

			u := tgbotapi.NewUpdate(0)
			u.Timeout = 60

			updates := bot.GetUpdatesChan(u)

			for update := range updates {

				if ctx.Err() == context.Canceled {
					break
				}

				if update.Message != nil { // If we got a message
					log.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					msg.ReplyToMessageID = update.Message.MessageID
					msg.ParseMode = "markdown"

					bot.Send(msg)
				}
			}
		},
	}
}
