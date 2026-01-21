package ai

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"google.golang.org/genai"
)

var (
	GEMINII_API string = os.Getenv("GEMINI_TOKEN")
	client      *genai.Client
	once        sync.Once
)

func initClient() {
	var err error
	ctx := context.Background()
	client, err = genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  GEMINII_API,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal("[ERROR] Gagal connect ke Gemini:", err)
	}
}

func Gemini(message string) (string, error) {
	once.Do(initClient)
	if client == nil {
		return "", fmt.Errorf("Client Gemini gagal di inisialisasi")
	}

	ctx := context.Background()
	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-3-flash-preview",
		genai.Text(message),
		nil,
	)

	if err != nil {
		return "", err
	}
	return result.Text(), nil
}
