package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

// Replace with your bot token from Discord Developer Portal
var token = "MTIwODM4MDQwOTgxNDE4ODA0Mg.GufPW0.G5hGQ93fQrlbTodsc7iL_9RKow7l8u2gEXazKQ"

// Function to handle when the bot is ready
func onReady(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Bot is now running. Press CTRL+C to exit.")
}

// Function to handle messages
func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Respond to a specific command
	if m.Content == "!hello" {
		s.ChannelMessageSend(m.ChannelID, "Hello, world!")
	}
}

func main() {
	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register message handler
	dg.AddMessageCreate(onMessage)
	dg.AddReadyHandler(onReady)

	// Open a websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	defer dg.Close()

	// Wait for the bot to be shut down (CTRL+C)
	fmt.Println("Bot is running. Press CTRL+C to exit.")
	select {}
}
