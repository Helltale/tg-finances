package handlers

import (
	"log"

	"github.com/helltale/tg-finances/client/internal/messages"
	"github.com/helltale/tg-finances/client/model"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func New(bot *telego.Bot, updates <-chan telego.Update, options ...th.BotHandlerOption) {

	bh, err := th.NewBotHandler(bot, updates)
	if err != nil {
		log.Fatalf("Ошибка при создании обработчика бота: %v", err)
	}
	defer bh.Stop()

	sessionContext := &model.SessionContext{
		State: model.StateStart,
	}

	bh.Handle(func(b *telego.Bot, update telego.Update) {
		if update.Message != nil {
			sessionContext.UserID = update.Message.From.ID

			log.Printf("`%d` send /start command", sessionContext.UserID)

			messages.SendMessageWithoutDelete(b, "hi hi hi", sessionContext)
		} else {
			log.Println("received update without message")
		}
	}, th.CommandEqual("start"))

	log.Println("handlers are ready")
	bh.Start()
}
