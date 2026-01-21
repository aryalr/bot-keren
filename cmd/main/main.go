package main

import (
	"bot-keren/pkg/dcbot"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ERROR] Error loading env:", err)
	}
	var BotToken = os.Getenv("BOT_TOKEN")
	dcbot.Run(BotToken)
}
