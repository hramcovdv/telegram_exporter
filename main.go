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
		log.Fatal("No telegram bot api token")
	}

	log.Print("Authorized on account ", bot.BotName())
	go bot.Run()

	srv := api.NewServer()

	log.Print("Server listen on ", listenAddr)
	log.Fatal(srv.Start(listenAddr))
}
