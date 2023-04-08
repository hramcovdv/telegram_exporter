package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/hramcovdv/telegram_exporter/api"
)

var (
	listenAddr string
)

func init() {
	flag.StringVar(&listenAddr, "listen", ":2112", "HTTP service listen address")
	flag.Parse()

	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	bot, err := api.NewBot(os.Getenv("TELEGRAM_BOT_API_TOKEN"))
	if err != nil {
		log.Fatalf("Telegram API token not found")
	}

	go bot.Run()

	srv := api.NewServer(bot)

	log.Printf("Server start and listen %s", listenAddr)
	log.Fatal(srv.Start(listenAddr))
}
