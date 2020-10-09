package main

import (
	"flag"
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

func init() {
	// TODO change this to be environment variables
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.StringVar(&cfgPath, "c", "config.yaml", "Config file")
	flag.Parse()
}

var token string
var cfgPath string

func main() {

	// Load config for app
	botCfg, err := config.New(cfgPath)

	if err != nil {
		log.Fatalf("Error whilst trying to read config file at %s: %v", cfgPath, err)
		return
	}

	// Create new log file for app
	f, err := logging.CreateLogFile(botCfg.Logging.Path, botCfg.App.Name)
	defer f.Close()

	// Open database session
	DBSession, err := bolt.Open(botCfg.DB.Name, botCfg.DB.Mode, nil)
	if err != nil {
		log.Fatalf("Error creating database session: %v", err)
		return
	}
	defer DBSession.Close()

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
	discord.Open()
	defer discord.Close()

	log.Println("Bot up and running!")
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	// Wait until signal is written to channel before killing bot
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Println("Kill signal received - exiting bot...")
}
