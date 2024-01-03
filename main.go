package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/germanpachec0/discord-bot/bot"
	"github.com/germanpachec0/discord-bot/config"
	"gopkg.in/yaml.v3"
)

func main() {
	file, err := os.ReadFile("config/creds.yml")
	if err != nil {
		log.Fatalf("Error reading yaml file: %v", err)
		return
	}
	var creds config.AppCredentials
	err = yaml.Unmarshal(file, &creds)
	if err != nil {
		log.Fatalf("Error reading yaml file: %v", err)
		return

	}
	session, err := bot.NewBot(creds)

	session.AddHandler(messageCreate)
	session.Identify.Intents = discordgo.IntentsGuildMessages

	err = session.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	fmt.Println("Bot is running, Press CTRL + C to exit")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if strings.TrimSpace(m.Content) == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if strings.TrimSpace(m.Content) == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
