package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/imakebank2/shortbot/birthdays"
	"github.com/imakebank2/shortbot/commands"
	"github.com/imakebank2/shortbot/eventhandlers"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file. Make you you create it from .env.example.")
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")

	if token == "" {
		log.Fatal("Paste the bot token (from Discord developer portal) into DISCORD_BOT_TOKEN.")
	}

	session, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	// Add event handlers here:
	session.AddHandler(eventhandlers.ResponseManager)
	session.AddHandler(eventhandlers.CommandManager)

	// Events bot can receive (should be the minimum possible for least overhead)
	session.Identify.Intents = discordgo.IntentsAll

	if err = session.Open(); err != nil {
		log.Fatal(err)
	}

	defer log.Println("Bot offline.")
	defer session.Close()

	commands.RegisterCommands(session)

	log.Println("Bot online! Press CTRL-C to exit.")

	log.Println("Birthdays calendar has started.")

	// Main chat
	// mainChat := "1490651667744555039"
	// session.ChannelMessageSend(mainChat, "Same")
	go birthdays.CheckBirthdays(session, "./birthdays/birthdays.json", "main-chat")

	// Blocks until CTRL-C or other signal is received
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}
