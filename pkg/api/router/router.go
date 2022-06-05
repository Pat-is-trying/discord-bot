package router

import (
	"context"
	"discord-bot/pkg/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var wg = sync.WaitGroup{}

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
		joke := fetchJoke(s, m)
		
		wg.Add(2)
		go s.ChannelMessageSend(m.ChannelID, joke.Joke)
		go saveJoke(joke)
		wg.Wait()
	}
}

func messageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {

	s.ChannelMessageSend(m.ChannelID, "https://stayhipp.com/wp-content/uploads/2019/02/you-better-watch.jpg")
}

type JokeResponse struct {
	Joke string `json:"joke" bson:"joke"`
}

func (j *JokeResponse) GetDoc() (primitive.D) {
	return bson.D{{"joke", j.Joke}}
}

func fetchJoke(s *discordgo.Session, m *discordgo.MessageCreate) (*JokeResponse) {

	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://icanhazdadjoke.com", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Bot Testing")

	res, _ := client.Do(req)

	j := &JokeResponse{}
	data, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(data, &j)
	if err != nil {
		log.Panic(err)
	}
	return j
}

func saveJoke(j database.DbConnected) {

	client := database.Connect()
	coll := client.Database(database.DATABASE).Collection(database.SCHEMA_JOKES)

	fmt.Println("Saving joke...")
	_, err := coll.InsertOne(context.TODO(), j.GetDoc())
	if err != nil {
		fmt.Println("Error saving joke!")
		return
	}
	fmt.Println("Joke Saved!")
}