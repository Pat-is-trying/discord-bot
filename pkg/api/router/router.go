package router

import (
	"github.com/bwmarrin/discordgo"
)

func SetRouter(s *discordgo.Session) {
	s.AddHandler(messageCreate)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	
	if m.Author.Bot {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	} else if m.Content == "bing" {
		s.ChannelMessageSend(m.ChannelID, "Bonged ya idiot!")
	} else if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}