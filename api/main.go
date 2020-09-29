package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"gopkg.in/yaml.v2"

	"github.com/aerzz23/visadiscordbot/api/handlers"
	"github.com/aerzz23/visadiscordbot/api/logging"
	"github.com/boltdb/bolt"
	"github.com/bwmarrin/discordgo"
)

type botConfig struct {
	DB struct {
		Name string      `yaml:"name"`
		Mode os.FileMode `yaml:"mode"`
	} `yaml:"db"`
	Logging struct {
		Path string `yaml:"path"`
	} `yaml:"logging"`
	App struct {
		Name string `yaml:"name"`
	} `yaml:"app"`
}

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.StringVar(&cfgPath, "c", "config.yaml", "Config file")
	flag.Parse()
}

var token string
var cfgPath string

// BotConfig is the config for app.
var BotConfig *botConfig

// DBSession is the BoltDB session for the app.
var DBSession *bolt.DB

func main() {

	// Load config for app
	BotConfig, err := newConfig(cfgPath)

	if err != nil {
		log.Fatalf("Error whilst trying to read config file at %s: %v", cfgPath, err)
		return
	}

	// Create new log file for app
	f, err := logging.CreateLogFile(BotConfig.Logging.Path, BotConfig.App.Name)
	defer f.Close()

	// Open database session
	DBSession, err := bolt.Open(BotConfig.DB.Name, BotConfig.DB.Mode, nil)
	if err != nil {
		log.Fatalf("Error creating database session: %v", err)
		return
	}
	defer DBSession.Close()

	botHandlers := handlers.New(DBSession)

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

func newConfig(cfgPath string) (*botConfig, error) {
	cfg := &botConfig{}

	f, err := os.Open(cfgPath)

	if err != nil {
		log.Fatalf("Error opening config file sat %s: %v", cfgPath, err)
		return nil, err
	}

	defer f.Close()

	d := yaml.NewDecoder(f)

	if err := d.Decode(&cfg); err != nil {
		log.Fatalf("Error decoding config file: %v", err)
		return nil, err
	}

	return cfg, nil
}
