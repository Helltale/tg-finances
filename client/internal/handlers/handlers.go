package handlers

import (
	"fmt"
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

			message := fmt.Sprintf(`Привет, %s!
Ближайшая цель: ....
Расход менльше на 10%%
Последняя потраченная сумма: ....`, update.Message.From.FirstName)

			messages.SendMainMenu(bot, message, sessionContext)
		} else {
			log.Println("received update without message")
		}
	}, th.CommandEqual("start"))

	log.Println("handlers are ready")
	bh.Start()
}
