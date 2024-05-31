# Gemini Discord Bot

A Discord bot built in Go that uses Gemini's 1.5 Flash model to provide conversational AI capabilities.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Bot](#running-the-bot)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- Conversational AI using Gemini's 1.5 Flash model
- Simple setup and configuration
- Real-time interaction on Discord servers

## Requirements

- Go 1.22.2 or higher
- A Discord Bot token
- Gemini API key

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/R894/gemini-discord-bot.git
    cd gemini-discord-bot
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

## Configuration

1. Create a `.env` file in the root directory and add your Discord Bot token and Gemini API key:

    ```env
    BOT_TOKEN=your-discord-bot-token
    API_KEY=your-gemini-api-key
    ```

## Running the Bot

To start the bot, simply run:

    ```sh
    go run main.go
    ```

Or build the project and run the executable:

    ```sh
    go build -o gemini-discord-bot
    ./gemini-discord-bot
    ```

## Usage

Once the bot is running, you can interact with it on Discord. Use the configured prefix (default `!`) to send commands.

Example:

    !gen Hello, how are you?


The bot will respond using Gemini's 1.5 model.

## License

Distributed under the MIT License. See `LICENSE` for more information.