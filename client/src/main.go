package main

import (
	"fmt"
	"log"

	"github.com/helltale/tg-finances/client/config"
	"github.com/helltale/tg-finances/client/internal/handlers"
	"github.com/helltale/tg-finances/client/internal/inlinehelper"
	"github.com/mymmrac/telego"
)

func main() {
	config, err := config.Get()
	if err != nil {
		log.Fatalf("error: bad token: %s", err)
	}

	bot, err := telego.NewBot(config.Token, telego.WithDefaultLogger(false, true))
	if err != nil {
		log.Fatalf("error: can not start bot: %s", err)
	}

	botInfo, _ := bot.GetMe()

	fistLogMessage := fmt.Sprintf(
		`
============================== lauch ==============================
token:		%s

botId: 		%d 
name: 		%s
isPremium: 	%t

author: 	github.com/helltale
===================================================================
`,
		config.Token,
		botInfo.ID,
		botInfo.FirstName,
		botInfo.IsPremium,
	)

	fmt.Println(fistLogMessage)

	inlinehelper.New(bot)

	updates, err := bot.UpdatesViaLongPolling(nil)
	if err != nil {
		fmt.Printf("error: error long polling: %s\n", err)
	}
	defer bot.StopLongPolling()

	handlers.New(bot, updates)

}
