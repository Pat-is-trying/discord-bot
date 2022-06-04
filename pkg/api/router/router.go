package router

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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
	} else if m.Content == "joke" {
		var joke *JokeResponse = &JokeResponse{}
		joke.getJoke(s, m)
	}
}

func messageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {

	s.ChannelMessageSend(m.ChannelID, "https://stayhipp.com/wp-content/uploads/2019/02/you-better-watch.jpg")
}

// Playing around with structs
type JokeResponse struct {
	Joke string `json:"joke"`
}

// Playing around with methods
func (j *JokeResponse) getJoke(s *discordgo.Session, m *discordgo.MessageCreate) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://icanhazdadjoke.com", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Bot Testing")

	res, _ := client.Do(req)

	data, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(data, &j)
	if err != nil {
		log.Panic(err)
	}

	s.ChannelMessageSend(m.ChannelID, j.Joke)
}