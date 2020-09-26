package main

import (
	"flag"
	"log"

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
	f, err := CreateLogFile(logPath, appName)
	discord, err := discordgo.New("Bot" + token)

	if err != nil {
		log.Println("Error creating Discord session - please check Token")
		return
	}

	log.Println(discord.Token)

	// discord.AddHandler()
}
