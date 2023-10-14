package main

import (
	"fmt"
	"log"
	"notifier/internal/server"
	"notifier/internal/telegram"
	"os"
)

func main() {
	log.Println("Starting...")
	tg_token := os.Getenv("TG_TOKEN")
	log.Println("Token: ", tg_token)
	// telegramClient := telegram.Login(tg_token)
	telegramClient := telegram.Client{}
	fmt.Println("1")
	telegramClient.Login(tg_token)
	httpServer := server.NewServer(&telegramClient)
	httpServer.Run(":8080")
}
