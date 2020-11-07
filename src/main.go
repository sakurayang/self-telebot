package main

import (
	"core/bot"
	"utils"
)

func main() {
	config := utils.GetConfig()
	bot.Init(config)
}
