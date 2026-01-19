package main

import (
	"bot-keren/pkg/dcbot"
)

// Token tetap di main atau bisa dipindah ke environment variable nanti
const BotToken = "TOKEN_BARU_ANDA_DISINI"

func main() {
	// Satu baris untuk menjalankan semuanya!
	dcbot.Run(BotToken)
}
