package handlers

import (
	"fmt"
	"github.com/aerzz23/visadiscordbot/api/netflix"
	"log"
	"strconv"
	"strings"

	"github.com/aerzz23/visadiscordbot/api/config"
	"github.com/aerzz23/visadiscordbot/api/events"
	"github.com/boltdb/bolt"
	"github.com/bwmarrin/discordgo"
)

const (
	newEventCmd   = "NewEvent"
	showEventsCmd = "ShowEvents"
	newNetflixSuggestionCmd = "NewNetflixSuggestion"
	generalErrMsg = "Oops <@%s> - something wrong happened! Please try again."
	logoKey       = "embedLogo"
)

// MessageCreateHandler is a handler for when Discord messages are sent within the channel.
func (bh *BotHandlersImpl) MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore messages made by the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	var mentioned = false

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
	// TODO put code into separate files or funcs at least as handler is getting complex and should just be switch
	// TODO replace with switch statement
	// Respond to Hi with a Hi back to Author
	if mentioned && strings.Contains(m.Content, "Hi") {
		log.Println("Hi Message Received")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hi, %s!", m.Author))
		return
	}

	// TODO Change how command works - needs to be separated using commas or something for multi string inputs
	// Add a new event to DB if user uses NewEvent command.
	if mentioned && strings.Contains(m.Content, newEventCmd) {
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

		err := createDiscordEvent(bh.cfg, bh.db, inputSlice)

		if err != nil {
			log.Println(fmt.Sprintf("Error creating Discord event: %v", err))
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(generalErrMsg, m.Author.ID))
			return
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("OK. <@%s> - Your event has been added!", m.Author.ID))
		return
	}

	if mentioned && strings.Contains(m.Content, showEventsCmd) {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Sure <@%s>, getting events...", m.Author.ID))
		allEvents, err := events.GetAll(bh.cfg, bh.db)

		if err != nil {
			log.Println(fmt.Sprintf("Error getting all events from DB. Error: %v", err))
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(generalErrMsg, m.Author.ID))
			return
		}

		eventsAsTable := events.FormatAsTable(allEvents)

		embed := discordgo.MessageEmbed{
			Title: "Visa Games Events",
			Type:  discordgo.EmbedTypeRich,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Table",
					Value: eventsAsTable,
				},
			},
		}

		s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		return
	}


	// Add a new Netflix suggestion to DB if user uses NewNetflixSuggestion command.
	if mentioned && strings.Contains(m.Content, newNetflixSuggestionCmd) {
		log.Println(fmt.Sprintf("NewNetflixSuggestion triggered by %s. Message: %s", m.Author, m.Content))
		suggestions := strings.Fields(m.Content)[2:]
		err := updateNetflixSuggestions(m.Author, bh.db, suggestions)
		if err != nil {
			log.Println(fmt.Printf("Error updating netflix suggestions. Error: %v", err))
			return
		}
	}

}

// createDiscordEvent takes string inputs, makes a new Event type with them, and then calls Create function to add to DB.
func createDiscordEvent(cfg *config.BotConfig, db *bolt.DB, inputs []string) error {
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

	return event.Create(cfg, db)
}

func updateNetflixSuggestions(author *discordgo.User, db *bolt.DB, suggestions []string) error {
	for _, suggestion := range suggestions {
		// Here I would check the availability...IF THERE WAS A FREE API
		ns := netflix.New(suggestion, author.Username)
		err := ns.Create(db, author.Username)
		if err != nil {
			log.Println(fmt.Printf("Error updating Netflix Suggestions. Error: %v", err))
			return err
		}
	}
	return nil
}
