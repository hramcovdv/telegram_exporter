package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/hramcovdv/telegram_exporter/api"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	bot, err := api.NewBot(os.Getenv("TELEGRAM_BOT_API_TOKEN"))
	if err != nil {
		log.Fatalf("Telegram API token not found")
	}

	srv := api.NewServer(bot)

	go bot.Run()

	log.Fatal(srv.Start(":3000"))
}
