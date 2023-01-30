package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var bot *tgbotapi.BotAPI

func sendTelegramMessage(msg string) {
	var err error
	tgToken := ""    // 要加入這個
	var chatID int64 // 要加入這個
	NewMsg := tgbotapi.NewMessage(chatID, msg)
	bot, err = tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		panic(err)
	}
	bot.Debug = false
	NewMsg.ParseMode = tgbotapi.ModeHTML //傳送html格式的訊息
	_, err = bot.Send(NewMsg)
	if err != nil {
		panic(err)
	}
}
