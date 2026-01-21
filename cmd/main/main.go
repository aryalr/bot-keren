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
		log.Println("[INFO] .env file not found, using system environment variables instead.", err)
	}
	var BotToken = os.Getenv("BOT_TOKEN")
	if BotToken == "" {
		log.Fatal("[ERROR] BOT_TOKEN is empty! Please set it in Railway variables.")
	}
	dcbot.Run(BotToken)
}
