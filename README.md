# VisaDiscordBot
VisaDiscordBot is a Discord Bot written in Go for the Visa Discord Channel. 

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine. 

See also instructions for Docker Compose for local MySQL instance.
### Prerequisites
* [Go 1.15](https://golang.org/dl/)
* [Docker](https://www.docker.com/products/docker-desktop)

### Installing
```
git clone https://github.com/Aerzz23/visadiscordbot.git
cd visadiscordbot
go run api/main.go -t "your_discord_api_token" -n "your_app_name" -l "your_log_location"
```

## Built With
* [Go 1.15](https://golang.org/doc/) - Programming Language
* [Docker](https://docs.docker.com/) - Containerization
* [discordgo](https://github.com/bwmarrin/discordgo) - Go bindings for Discord

## Authors
* **Aaron Burge** - *Creator* - [aerzz23](https://github.com/Aerzz23)