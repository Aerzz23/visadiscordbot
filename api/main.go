package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aerzz23/visadiscordbot/api/handlers"
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

	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Println("Error creating Discord session - please check Token")
		return
	}

	log.Println(discord.Token)

	log.Println("Adding handler for MessageCreate")
	discord.AddHandler(handlers.MessageCreateHandler)

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
