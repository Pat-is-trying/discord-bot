package router

import (
	"github.com/bwmarrin/discordgo"
)

func SetRouter(s *discordgo.Session) {
	s.AddHandler(messageCreate)
	s.AddHandler(messageDelete)
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

func messageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {

	s.ChannelMessageSend(m.ChannelID, "https://stayhipp.com/wp-content/uploads/2019/02/you-better-watch.jpg")
}