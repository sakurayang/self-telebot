package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"utils"
)

func Init(config utils.Config) *tgbotapi.BotAPI {
	token := config.Bot.Token
	apiHost := config.Bot.ApiHost
	var bot *tgbotapi.BotAPI
	var err error
	if config.Proxy.Enable {
		host := config.Proxy.Host
		port := strconv.Itoa(config.Proxy.Port)
		proxyUrl, _ := url.Parse(host + ":" + port)
		client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
		bot, err = tgbotapi.NewBotAPIWithClient(token, apiHost, client)
	} else {
		bot, err = tgbotapi.NewBotAPIWithAPIEndpoint(token, apiHost)
	}
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = config.Bot.Debug

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		go func(update tgbotapi.Update) {
			msg := parseMessage(update.Message)
			Send(bot, update.Message.Chat.ID, msg)
		}(update)
	}

	return bot
}

func Send(bot *tgbotapi.BotAPI, chatId int64, msg string) {
	send_msg := tgbotapi.NewMessage(chatId, msg)
	go bot.Send(send_msg)
}

func parseMessage(message *tgbotapi.Message) string {
	msg := ""
	if message.IsCommand() {
		msg = invokeCommand(message)
	}
	return msg
}

func invokeCommand(message *tgbotapi.Message) string {
	command := message.Command()
	msg := ""
	switch command {
	case "help":
		msg = "/help - show help"
	}
	return msg
}
