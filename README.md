# VisaDiscordBot

VisaDiscordBot is a Discord Bot written in Go for the Visa Discord Channel.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine.

Additonal instructions are also provided for Docker.

### Prerequisites

* [Go 1.15](https://golang.org/dl/)
* [Docker](https://www.docker.com/products/docker-desktop)

### Running

#### Go Local

```bash
git clone https://github.com/Aerzz23/visadiscordbot.git
cd visadiscordbot
export VISA_BOT_CONFIG="YOUR_CONFIG_PATH_HERE" # absolute path
export VISA_BOT_TOKEN="YOUR_DISCORD_BOT_TOKEN_HERE" # https://discord.com/developers/applications
go get -u
go run api/main.go

```

#### Docker

```bash
git clone https://github.com/Aerzz23/visadiscordbot.git
cd visadiscordbot
docker build -t aerzz23/visadiscordbot:latest .
docker run --env VISA_BOT_TOKEN=YOUR_DISCORD_BOT_TOKEN_HERE aerzz23/visadiscordbot # https://discord.com/developers/

```

### Testing

Unit Tests are written using the [Ginkgo](https://github.com/onsi/ginkgo) framework.

```bash
cd visadiscordbot
go test ./...

```

#### Or

```bash

cd visadiscordbot
ginkgo ./...

```

## Built With

* [Go 1.15](https://golang.org/doc/) - Programming Language
* [DiscordGo](https://github.com/bwmarrin/discordgo) - Go bindings for Discord
* [ASCII Table Writer](https://github.com/olekukonko/tablewriter) - Go library for  ASCII formatted tables
* [BoltDB](https://github.com/boltdb/bolt) - Database
* [Ginkgo](https://github.com/onsi/ginkgo) - BDD Testing Framework
* [Docker](https://docs.docker.com/) - Containerization

## Authors

* **Aaron Burge** - *Creator* - [Aerzz23](https://github.com/Aerzz23)
