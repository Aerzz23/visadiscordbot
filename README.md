# VisaDiscordBot
VisaDiscordBot is a Discord Bot written in Go for the Visa Discord Channel. 

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine. 

Additonal instructions are also provided for Docker.
### Prerequisites
* [Go 1.15](https://golang.org/dl/)
* [Docker](https://www.docker.com/products/docker-desktop)

### Installing
```
git clone https://github.com/Aerzz23/visadiscordbot.git
cd visadiscordbot
go get -u 
go run api/main.go -t "your_discord_api_token" -c "config_path"
```

## Built With
* [Go 1.15](https://golang.org/doc/) - Programming Language
* [BoltDB](https://github.com/boltdb/bolt) - Database
* [Docker](https://docs.docker.com/) - Containerization
* [discordgo](https://github.com/bwmarrin/discordgo) - Go bindings for Discord

## Authors
* **Aaron Burge** - *Creator* - [aerzz23](https://github.com/Aerzz23)