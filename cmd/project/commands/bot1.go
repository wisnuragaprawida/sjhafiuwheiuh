package commands

import (
	"context"
	"encoding/json"
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

			b, err := bot.New("6300175469:AAHly94Qz_4pdZIZ3_x06Y1GXt7g-g-_6Ug", opts...)
			// b, err := bot.New("5126877034:AAFkUGUuS6d-z6TDLe6NIWgLcUUGi73U3to", opts...)
			if err != nil {
				log.Error(err)
				return
			}
			b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, myStartHandler)

			log.Info("Bot started")

			b.Start(ctx)
		},
	}
}

func myStartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Info("isi nya ", update.Message.Chat.ID)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Hello, World!",
	})
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {

	shal, err := json.Marshal(update)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(string(shal))

	// log.Info("isi nya ", update.ChannelPost.Text)

	// b.ForwardMessage(ctx, &bot.ForwardMessageParams{})

	// if update.ChannelPost.Text == "fwd" {
	// 	b.ForwardMessage(ctx, &bot.ForwardMessageParams{
	// 		ChatID:     1288114342,
	// 		FromChatID: "-838472573",
	// 		MessageID:  275,
	// 	})

	// }

	// // photo :=
	if update.ChannelPost.Text == "del" {
		_, errs := b.SendPhoto(ctx, &bot.SendPhotoParams{
			ChatID: 1288114342,
			//from chat id

			Photo:   &models.InputFileString{Data: "AgACAgUAAx0CdZ6c1wADDmSsPTGkEzE1Yu7yHmBOp9gKP7N5AAJ6ujEbQ5tgVfQNZLYVNhDoAQADAgADcwADLwQ"},
			Caption: "test",
		})
		if errs != nil {
			log.Error(errs)
		}
	}
	//we can do whatever we want with the update here

	// if update.Message.Text == "kon" {
	// 	b.SendMessage(ctx, &bot.SendMessageParams{
	// 		ChatID:    update.Message.Chat.ID,
	// 		Text:      "tol",
	// 		ParseMode: models.ParseModeHTML,
	// 	})
	// }

}
