package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/joho/godotenv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func sendMediaGroup(bot *tgbotapi.BotAPI, chatID int64, imageURL string) {
	photo := tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(imageURL))
	mediaGroup := tgbotapi.NewMediaGroup(chatID, []interface{}{photo})

	messages, err := bot.SendMediaGroup(mediaGroup)
	if err != nil {
		log.Printf("Error sending media group: %v", err)
		return
	}

	for _, msg := range messages {
		log.Printf("Sent message: %v", msg)
	}
}

func getRandomImage() string {
	randomNumber := rand.Intn(1000) + 1
	return fmt.Sprintf("images/%03d.png", randomNumber)
}

// Notify the second bot about the user interaction
func notifyUsage(userID int64, username string, command string) {
	YOUR_PERSONAL_CHAT_ID_STRING := os.Getenv("YOUR_PERSONAL_CHAT_ID")
	YOUR_PERSONAL_CHAT_ID, err := strconv.ParseInt(YOUR_PERSONAL_CHAT_ID_STRING, 10, 64)
	if err != nil {
		log.Fatalf("Error parsing YOUR_PERSONAL_CHAT_ID: %v", err)
	}

	apiToken := os.Getenv("API_BOT_TOKEN")
	if apiToken == "" {
		log.Fatal("API_BOT_TOKEN environment variable is required")
	}

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", apiToken)
	messageText := fmt.Sprintf("User @%s (%d) used command: %s", username, userID, command)
	data := url.Values{
		"chat_id": {strconv.FormatInt(YOUR_PERSONAL_CHAT_ID, 10)},
		"text":    {messageText},
	}

	_, err = http.PostForm(apiURL, data)
	if err != nil {
		log.Printf("Error notifying usage: %v", err)
	}
}

func main() {


	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is required")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	rand.Seed(time.Now().UnixNano())

	for update := range updates {
		if update.Message == nil {
			continue
		}

		userID := update.Message.From.ID
		username := update.Message.From.UserName

		// Use a goroutine to run notifyUsage concurrently
		go notifyUsage(userID, username, update.Message.Text)

		switch {
		case strings.HasPrefix(update.Message.Text, "/start"):
			welcomeMsg := "🌟 Welcome to Just Little Things! 🌟\n\nHere, we celebrate the small joys that make life beautiful. Use the following commands:\n\n/littlething - Get a random image\n/littlething [number] - Get a specific image (1-1000)\n/discoverjoy - Another way to get a random image\n/help - See available commands\n\nEnjoy exploring the simple pleasures!"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcomeMsg)
			bot.Send(msg)

		case strings.HasPrefix(update.Message.Text, "/littlething"):
			args := strings.Fields(update.Message.Text)
			if len(args) > 1 {
				num, err := strconv.Atoi(args[1])
				if err == nil && num >= 1 && num <= 1000 {
					imageURL := fmt.Sprintf("images/%03d.png", num)
					sendMediaGroup(bot, update.Message.Chat.ID, imageURL)
				} else {
					sendMediaGroup(bot, update.Message.Chat.ID, getRandomImage())
				}
			} else {
				sendMediaGroup(bot, update.Message.Chat.ID, getRandomImage())
			}

		case update.Message.Text == "/discoverjoy":
			sendMediaGroup(bot, update.Message.Chat.ID, getRandomImage())

		case update.Message.Text == "/help":
			helpMsg := "🛠 Available commands: 🛠\n\n/start - Welcome message and instructions\n/littlething - Get a random image\n/littlething [number] - Get a specific image (1-1000)\n/discoverjoy - Another way to get a random image\n/help - See this list of commands"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpMsg)
			bot.Send(msg)

		default:
			unknownMsg := "🤔 Hmm, I didn't quite catch that. Type /help to see what I can do for you!"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, unknownMsg)
			bot.Send(msg)
		}
	}
}
