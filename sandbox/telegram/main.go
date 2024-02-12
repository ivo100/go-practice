package main

import (
	"log"
	"os"

	tel "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Panic("TELEGRAM_TOKEN is not set")
	}
	bot, err := tel.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	//chatId := readUpdates(bot)
	//username := "@ivostoy"
	chatID := int64(6147574575)
	//user, err := bot.GetChat(tel.ChatConfig{SuperGroupUsername: username})
	//if err != nil {
	//	log.Panic(err)
	//}
	//chatID := user.ID
	//chatID, err := getChatID(bot, recipientUsername)
	//if err != nil {
	//	log.Panic(err)
	//}
	// Create a new message configuration
	msg := tel.NewMessage(chatID, "ALERT ALERT ALERT!")

	// Optionally, you can customize the message settings here
	// For example, to disable notifications for this message:
	// msg.DisableNotification = true

	// Send the message using the bot API
	_, err = bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Message sent successfully!")
}

func readUpdates(bot *tel.BotAPI) int64 {
	chatID := int64(6147574575)
	// Set up a webhook to receive updates (optional)
	// You can alternatively use polling to receive updates
	// bot.RemoveWebhook()
	// Get updates from the Telegram Bot API
	updateConfig := tel.NewUpdate(0)
	updateConfig.Timeout = 60
	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Panic(err)
	}
	for update := range updates {
		if update.Message != nil {
			chatID = update.Message.Chat.ID
			log.Printf("Chat ID: %d", chatID)
			break // Exit after printing the first chat ID
		}
	}
	return chatID
}
