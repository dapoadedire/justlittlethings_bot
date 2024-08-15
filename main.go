package main

import (
    "fmt"
    "log"
    "math/rand"
    "strconv"
    "strings"
    "time"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func sendMediaGroup(bot *tgbotapi.BotAPI, chatID int64, imageURL string) {
    photo := tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(imageURL))
    mediaGroup := tgbotapi.NewMediaGroup(chatID, []interface{}{photo})

    // Send the photo
    messages, err := bot.SendMediaGroup(mediaGroup)
    if err != nil {
        log.Panic(err)
    }

    // Handle the response messages if needed
    for _, msg := range messages {
        log.Printf("Sent message: %v", msg)
    }
}

func main() {
    // Replace with your Telegram bot token
    botToken := "7222244368:AAGEINzLHL6I2rb_dIaLHhb23ONzTKxb_Ng"

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
        if update.Message == nil { // ignore any non-Message updates
            continue
        }

        switch {
        case strings.HasPrefix(update.Message.Text, "/littlething"):
            args := strings.Fields(update.Message.Text)
            if len(args) > 1 {
                num, err := strconv.Atoi(args[1])
                if err == nil && num >= 1 && num <= 1000 {
                    imageURL := fmt.Sprintf("images/%03d.png", num)
                    sendMediaGroup(bot, update.Message.Chat.ID, imageURL)
                } else {
                    randomNumber := rand.Intn(1000) + 1
                    imageURL := fmt.Sprintf("images/%03d.png", randomNumber)
                    sendMediaGroup(bot, update.Message.Chat.ID, imageURL)
                }
            } else {
                randomNumber := rand.Intn(1000) + 1
                imageURL := fmt.Sprintf("images/%03d.png", randomNumber)
                sendMediaGroup(bot, update.Message.Chat.ID, imageURL)
            }
        case update.Message.Text == "/start":
            welcomeMsg := "🌟 Welcome to Just Little Things! 🌟\n\nHere, we celebrate the small joys that make life beautiful. To get started, type /littething [number] to view a specific delight (1-1000), or simply /littething for a random moment of happiness. Enjoy exploring the simple pleasures right under your nose!"
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcomeMsg)
            bot.Send(msg)
        case update.Message.Text == "/help":
            helpMsg := "🛠 Need some help? No worries! 🛠\n\nUse /littething [number] to view a specific image (1-1000) or simply /littething for a random surprise. Let's discover life's little joys together!"
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpMsg)
            bot.Send(msg)
        default:
            unknownMsg := "🤔 Hmm, I didn't quite catch that. Type /help to see what I can do for you!"
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, unknownMsg)
            bot.Send(msg)
        }
    }
}
