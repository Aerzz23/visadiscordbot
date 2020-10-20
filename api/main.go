package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aerzz23/visadiscordbot/api/config"
	"github.com/aerzz23/visadiscordbot/api/handlers"
	"github.com/aerzz23/visadiscordbot/api/logging"
	"github.com/boltdb/bolt"
	"github.com/bwmarrin/discordgo"
)

const (
	tokenEnv = "VISA_BOT_TOKEN"
	cfgEnv   = "VISA_BOT_CONFIG"
)

var (
	token   string
	cfgPath string
)

func init() {
	token = os.Getenv(tokenEnv)
	cfgPath = os.Getenv(cfgEnv)
}

func main() {

	// Load config for app
	botCfg, err := config.New(cfgPath)

	if err != nil {
		log.Fatalf("Error whilst trying to read config file at %s: %v", cfgPath, err)
		return
	}

	// Create new log file for app
	if _, err := os.Stat(botCfg.Logging.Path); os.IsNotExist(err) {
		err = os.Mkdir(botCfg.Logging.Path, os.ModePerm)
	}
	f, err := logging.CreateLogFile(botCfg.Logging.Path, botCfg.App.Name)
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			log.Fatal("Error closing log file")
			err = closeErr
		}
	}()

	// Open database session
	DBSession, err := bolt.Open(botCfg.DB.Name, botCfg.DB.Mode, nil)
	if err != nil {
		log.Fatalf("Error creating database session: %v", err)
		return
	}
	defer func() {
		if closeErr := DBSession.Close(); closeErr != nil {
			log.Fatal("Error closing DB session")
			err = closeErr
		}
	}()

	botHandlers := handlers.New(botCfg, DBSession)

	// Create new instance of discord go
	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Println("Error creating Discord session - please check Token")
		return
	}

	log.Println("Adding handler for MessageCreate")
	discord.AddHandler(botHandlers.MessageCreateHandler)

	// Open connection to Discord
	log.Println("Opening websocket to Discord")
	err = discord.Open()
	if err != nil {
		log.Fatalf("Error whilst trying to open websocket to Discord: %v", err)
		return
	}
	defer func() {
		if closeErr := discord.Close(); closeErr != nil {
			log.Fatal("Error closing Discord session")
			err = closeErr
		}
	}()

	log.Println("Bot up and running!")
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	// Wait until signal is written to channel before killing bot
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Println("Kill signal received - exiting bot...")
}
