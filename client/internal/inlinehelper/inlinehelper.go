package inlinehelper

import (
	"log"

	"github.com/mymmrac/telego"
)

func New(bot *telego.Bot) {
	commands := []telego.BotCommand{
		{Command: "start", Description: "Перезапустить бота"},
	}

	params := &telego.SetMyCommandsParams{
		Commands: commands,
	}

	err := bot.SetMyCommands(params)
	if err != nil {
		log.Fatalf("Ошибка при установке команд: %v", err)
	}

	log.Println("inline helper is ready")
}
