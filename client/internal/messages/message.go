package messages

import (
	"log"

	"github.com/helltale/tg-finances/client/model"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func SendMessageWithoutDelete(bot *telego.Bot, message string, sessionContext *model.SessionContext) {

	sentMessage, err := bot.SendMessage(
		tu.Message(
			tu.ID(sessionContext.UserID),
			message,
		),
	)
	if err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	} else {
		sessionContext.LastMessageIDs = append(sessionContext.LastMessageIDs, int64(sentMessage.MessageID))
	}
}
