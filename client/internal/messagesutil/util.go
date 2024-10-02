package messagesutil

import (
	"log"

	"github.com/helltale/tg-finances/client/model"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func DeleteMessages(bot *telego.Bot, sessionContext *model.SessionContext) {
	if len(sessionContext.LastMessageIDs) > 0 {
		for _, messageID := range sessionContext.LastMessageIDs {
			err := bot.DeleteMessage(&telego.DeleteMessageParams{
				ChatID:    tu.ID(sessionContext.UserID),
				MessageID: int(messageID),
			})
			if err != nil {
				log.Printf("error with deleting message: %v", err)
			}
		}

		sessionContext.LastMessageIDs = []int64{}
	}
}
