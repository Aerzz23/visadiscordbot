package handlers

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aerzz23/visadiscordbot/api/events"
	"github.com/boltdb/bolt"
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
		return
	}

	// Add a new event to DB if user uses NewEvent command.
	if mentioned && strings.Contains(m.Content, "NewEvent") {
		log.Println(fmt.Sprintf("NewEvent triggered by %s. Message: %s", m.Author, m.Content))
		inputSlice := strings.Fields(m.Content)[2:]

		numOfInputs := len(inputSlice)

		if numOfInputs != 4 && numOfInputs != 6 {
			log.Println(fmt.Sprintf("NewEvent has invalid amount of arguments, was given %s with %d arguments", inputSlice, numOfInputs))
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Oops <@%s> - it seems like you've given me the wrong amount of arguments. You gave me %d instead of 4 or 6", m.Author.ID, numOfInputs))
			return
		}

		if numOfInputs == 4 {
			log.Println(fmt.Sprintf("NewEvent triggered with no alert. Input: %s", m.Content))
			// Populate alert as false and alertTime as 0.
			inputSlice = append(inputSlice, "false", "0")
		}

		err := createDiscordEvent(bh.db, inputSlice)

		if err != nil {
			log.Println(fmt.Sprintf("Error creating Discord event: %v", err))
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Oops <@%s> - sommething wrong happened! Please try again.", m.Author.ID))
			return
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("OK. <@%s> - Your event has been added!", m.Author.ID))
		return
	}
}

// createDiscordEvent takes string inputs, makes a new Event type with them, and then calls Create function to add to DB.
func createDiscordEvent(db *bolt.DB, inputs []string) error {
	alert, err := strconv.ParseBool(inputs[4])
	if err != nil {
		log.Println(fmt.Sprintf("Error trying to parse alert as bool, was given %s. Error: %v", inputs[4], err))
		return err
	}

	alertTime, err := strconv.Atoi(inputs[5])

	if err != nil {
		log.Println(fmt.Sprintf("Error trying to parse alertTime as int, was given: %s. Error: %v", inputs[5], err))
		return err
	}

	event, err := events.New(inputs[0], inputs[1], inputs[2], inputs[3], alert, alertTime)

	if err != nil {
		log.Println(fmt.Printf("Error creating new event. Error: %v", err))
		return err
	}
	return event.Create(db)
}
