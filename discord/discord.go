package discord

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/R894/gemini-test/gemini"
	"github.com/bwmarrin/discordgo"
)

func Start() {
	dg, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	session := gemini.New()
	if err != nil {
		log.Fatal("error creating Discord session,", err)
	}
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		messageCreate(s, m, session)
	})

	dg.Identify.Intents = discordgo.IntentsGuildMessages
	err = dg.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	session.Close()
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate, gemini *gemini.Gemini) {

	// Ignore all messages created by the bot itself or other bots
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	if strings.Contains(m.Content, "!gen") {
		str, _ := strings.CutPrefix(m.Content, "!gen")

		content, err := gemini.SendMessage(fmt.Sprintf("%s: %s", m.Author.GlobalName, str))
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, string(content))
	}
}
