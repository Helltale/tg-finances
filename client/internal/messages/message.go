package messages

import (
	"log"

	"github.com/helltale/tg-finances/client/internal/messagesutil"
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

func SendMessage(bot *telego.Bot, message string, sessionContext *model.SessionContext) {
	messagesutil.DeleteMessages(bot, sessionContext)
	SendMessageWithoutDelete(bot, message, sessionContext)
}

func SendMessageInlineKeyboardWithoutDelete(bot *telego.Bot, messageText string, sessionContext *model.SessionContext, buttonText string, buttonCallbackData string) {

	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(buttonText).WithCallbackData(buttonCallbackData),
		),
	)

	message := tu.Message(
		tu.ID(sessionContext.UserID),
		messageText,
	).WithReplyMarkup(inlineKeyboard)

	sentMessage, err := bot.SendMessage(message)
	if err != nil {
		log.Printf("Ошибка при отправке сообщения (id: `%d` 	text: `%s`) : %v", sentMessage.MessageID, sentMessage.Text, err)
	} else {
		sessionContext.LastMessageIDs = append(sessionContext.LastMessageIDs, int64(sentMessage.MessageID))
	}
}

func SendMessageInlineKeyboard(bot *telego.Bot, messageText string, sessionContext *model.SessionContext, buttonText string, buttonCallbackData string) {
	messagesutil.DeleteMessages(bot, sessionContext)
	SendMessageInlineKeyboardWithoutDelete(bot, messageText, sessionContext, buttonText, buttonCallbackData)
}

func SendMainMenuWithoutDelete(bot *telego.Bot, messageText string, sessionContext *model.SessionContext) {
	inlineKeyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Расходы").WithCallbackData("teachers"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Доходы").WithCallbackData("contacts"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Цели").WithCallbackData("contacts"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Кэшбеки").WithCallbackData("contacts"),
		),
	)
	message := tu.Message(
		tu.ID(sessionContext.UserID),
		messageText,
	).WithReplyMarkup(inlineKeyboard)

	sentMessage, err := bot.SendMessage(message)
	if err != nil {
		log.Printf("Ошибка при отправке меню: %v", err)
	} else {
		sessionContext.LastMessageIDs = append(sessionContext.LastMessageIDs, int64(sentMessage.MessageID))
	}
}

func SendMainMenu(bot *telego.Bot, messageText string, sessionContext *model.SessionContext) {
	messagesutil.DeleteMessages(bot, sessionContext)
	SendMainMenuWithoutDelete(bot, messageText, sessionContext)
}
