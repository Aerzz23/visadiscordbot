package main

import (
	"flag"
	"log"

	"github.com/aerzz23/visadiscordbot/api/logging"
	"github.com/bwmarrin/discordgo"
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.StringVar(&appName, "n", "", "App Name")
	flag.StringVar(&logPath, "l", "", "Log Path")
	flag.Parse()
}

var token string
var appName string
var logPath string

func main() {
	f, err := logging.CreateLogFile(logPath, appName)
	defer f.Close()

	discord, err := discordgo.New("Bot" + token)

	if err != nil {
		log.Println("Error creating Discord session - please check Token")
		return
	}

	log.Println(discord.Token)

	// discord.AddHandler()
}
