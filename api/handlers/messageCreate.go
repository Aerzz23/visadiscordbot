package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// MessageCreateHandler is a handler for when Discord messages are sent within the channel.
func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore messages made by the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	var mentioned bool = false

	// Check to see if Bot was mentioned - otherwise ignore message
	if len(m.Mentions) > 0 {
		for _, user := range m.Mentions {
			if user.ID == s.State.User.ID {
				log.Println(fmt.Sprintf("Bot was mentioned by %s. Message: %s", m.Author, m.Content))
				mentioned = true
				break
			}
		}
	}

	if mentioned && strings.Contains(m.Content, "Hi") {
		log.Println("Hi Message Received")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hi, %s!", m.Author))
	}

}
