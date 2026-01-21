package dcbot

import (
	"bot-keren/pkg/ai"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Fungsi Run: Menangani inisialisasi, koneksi, dan blocking
func Run(token string) {
	// 1. Setup Session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error membuat sesi Discord:", err)
	}

	// 2. Register Handler
	dg.AddHandler(MessageCreate)

	// 3. Setup Intents
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent

	// 4. Open Connection
	err = dg.Open()
	if err != nil {
		log.Fatal("Error membuka koneksi WebSocket:", err)
	}

	fmt.Println("Bot sedang berjalan... Tekan CTRL-C untuk berhenti.")

	// 5. Wait / Blocking (Agar program tidak langsung mati)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc // Program diam di sini sampai user menekan CTRL-C

	// 6. Clean Shutdown
	dg.Close()
	fmt.Println("\nBot dimatikan.")
}

// Message Handler
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	pesan := m.Content
	fmt.Printf("[LOG] User: %s | Pesan: %s\n", m.Author.Username, pesan)

	s.ChannelTyping(m.ChannelID)

	// Panggil AI Package
	balas, err := ai.Gemini(pesan)
	if err != nil {
		fmt.Println("[ERROR] Gemini:", err)
		s.ChannelMessageSend(m.ChannelID, "Maaf, engine sedang ada massalah (ERROR)")
		return
	}

	// Logika Split Pesan Panjang (>2000 char)
	if len(balas) > 2000 {
		limit := 1900
		runes := []rune(balas)
		for i := 0; i < len(runes); i += limit {
			end := i + limit
			if end > len(runes) {
				end = len(runes)
			}
			s.ChannelMessageSend(m.ChannelID, string(runes[i:end]))
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, balas)
	}
}
