package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// MessageCreateHandler is a handler for when Discord messages are sent within the channel.
func (bh *BotHandlersImpl) MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore messages made by the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	var mentioned bool = false

	// Check to see if Bot was mentioned - otherwise ignore message (might still other things where message created without mention TBD)
	if len(m.Mentions) > 0 {
		for _, user := range m.Mentions {
			if user.ID == s.State.User.ID {
				log.Println(fmt.Sprintf("Bot was mentioned by %s. Message: %s", m.Author, m.Content))
				mentioned = true
				break
			}
		}
	}

	// Respond to Hi with a Hi back to Author
	if mentioned && strings.Contains(m.Content, "Hi") {
		log.Println("Hi Message Received")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hi, %s!", m.Author))
	}

	if mentioned && strings.Contains(m.Content, "NewEvent") {
		log.Println(fmt.Sprintf("NewEvent triggered by %s. Message: %s", m.Author, m.Content))
		inputSlice := strings.Fields(m.Content)[2:]

		if lengthOfSlice := len(inputSlice); lengthOfSlice < 4 || lengthOfSlice > 6 {
			log.Println(fmt.Sprintf("NewEvent has invalid amount of arguments, was given %s with %d arguments", inputSlice, lengthOfSlice))
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Oops <@%s> - it seems like you've given me the wrong amount of arguments. You gave me %d instead of 4 or 6", m.Author.ID, lengthOfSlice))
			return
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("OK. <@%s> - Your event has been added!", m.Author.ID))

	}

}
