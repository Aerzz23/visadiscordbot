package handlers

import (
	"github.com/aerzz23/visadiscordbot/api/config"
	"github.com/boltdb/bolt"
	"github.com/bwmarrin/discordgo"
)

// BotHandlers is an interface for all handlers for the Bot.
// Examples of a Handler would be listening for a message creation.
type BotHandlers interface {
	MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate)
}

// BotHandlersImpl is a struct which has all handlers as receivers and stores state for handlers.
type BotHandlersImpl struct {
	cfg *config.BotConfig
	db  *bolt.DB
}

// New creates a new instance of BotHandlersImpl, storing BoltDB session.
func New(cfg *config.BotConfig, db *bolt.DB) *BotHandlersImpl {
	return &BotHandlersImpl{cfg: cfg, db: db}
}
