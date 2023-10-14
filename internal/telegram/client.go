package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client struct {
	Token   string
	Api     *tgbotapi.BotAPI
	chat_id int
}

// Login to telegram using API Token
func (c *Client) Login(apiToken string) *Client {
	api, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		log.Println("Unable to login to Telegram")
		log.Panic(err)
	}
	c.Api = api
	return &Client{}
}

func (c *Client) SendMessage(message string) error {
	chat_id := 189220134
	msg := tgbotapi.NewMessage(int64(chat_id), "Oh, Hi!")
	_, err := c.Api.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

// This is an example of how to use bot to handle events
func example() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
