package discord

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

type Discord struct {
	Token string
}

func (d *Discord) Init(db *gorm.DB) {
	client, err := discordgo.New("Bot " + d.Token)
	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	err = client.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return
	}

	log.Println("Bot is now running. " + client.State.User.Username)

	client.AddHandler(CreateMessage)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	client.Close()

}
