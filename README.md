# Just Little Things Bot

A Telegram bot that celebrates the small joys in life by sharing random or specific images.

## Description

Just Little Things Bot is a Telegram bot that allows users to receive images representing simple pleasures and little joys in life. Users can request random images or specify a particular image number.

GitHub Repository: [https://github.com/dapoadedire/justlittlethings_bot](https://github.com/dapoadedire/justlittlethings_bot)

## Features

- Get a random image of a little joy
- Request a specific image by number
- Simple and intuitive commands

## Commands

- `/start` - Welcome message and instructions
- `/littlething` - Get a random image
- `/littlething [number]` - Get a specific image (1-1000)
- `/discoverjoy` - Another way to get a random image
- `/help` - See available commands

## How to Run

1. Clone the repository:

```bash
git clone https://github.com/dapoadedire/justlittlethings_bot.git
cd justlittlethings_bot
```

2. Install dependencies:

```bash
go mod tidy
```

3. Set up environment variables:
   Create a `.env` file in the project root and add the following:

```bash
TELEGRAM_BOT_TOKEN=your_telegram_bot_token
API_BOT_TOKEN=your_api_bot_token
YOUR_PERSONAL_CHAT_ID=your_personal_chat_id
```

4. Ensure you have a folder named `images` in the project root containing numbered images from 001.png to 1000.png.

5. Run the bot:

```bash
go run main.go
```

6. Start chatting with the bot on Telegram.

## Dependencies

- github.com/joho/godotenv
- github.com/go-telegram-bot-api/telegram-bot-api/v5

## Contributing

Contributions, issues, and feature requests are welcome. Feel free to check the [issues page](https://github.com/dapoadedire/justlittlethings_bot/issues) if you want to contribute.

## License

This project is licensed under the MIT License.
