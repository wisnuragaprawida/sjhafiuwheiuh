package commands

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/spf13/cobra"
	"github.com/wisnuragaprawida/project/bootstrap"
	"github.com/wisnuragaprawida/project/pkg/log"
)

func init() {
	registerCommand(startBot1)
}

func startBot1(dep *bootstrap.Dependency) *cobra.Command {
	return &cobra.Command{
		Use:   "bot1",
		Short: "bot1 service",
		Run: func(cmd *cobra.Command, args []string) {
			dep.GetConfig()

			ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

			opts := []bot.Option{
				bot.WithDefaultHandler(handler),
			}

			bot, err := bot.New("5126877034:AAFkUGUuS6d-z6TDLe6NIWgLcUUGi73U3to", opts...)
			if err != nil {
				log.Error(err)
				return
			}

			bot.Start(ctx)
		},
	}
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
