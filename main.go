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
        case strings.HasPrefix(update.Message.Text, "/sendpic"):
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
            welcomeMsg := "Welcome! Send /sendpic [number] to receive a specific image or just /sendpic for a random image."
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcomeMsg)
            bot.Send(msg)
        case update.Message.Text == "/help":
            helpMsg := "Use /sendpic [number] to get a specific image (1-1000) or just /sendpic for a random image."
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpMsg)
            bot.Send(msg)
        default:
            unknownMsg := "Unknown command. Type /help for available commands."
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, unknownMsg)
            bot.Send(msg)
        }
    }
}
