package session

import (
	"discord-bot/pkg/auth"
	"log"

	"github.com/bwmarrin/discordgo"
)

func InitSession() (*discordgo.Session) {
	token := auth.FetchEnv("TOKEN")

	// Create a new Discord session using the provided bot token.
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("error creating Discord session,", err)
	}
	
	return s
}