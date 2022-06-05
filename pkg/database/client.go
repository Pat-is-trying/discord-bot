package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb+srv://ppickard:Deadmau5@bot-cluster.sg81c.mongodb.net/?retryWrites=true&w=majority"
const DATABASE = "discord-bot"
const SCHEMA_JOKES = "jokes"

// Connects to the MongoDB and initializes the client object.
func Connect() (*mongo.Client) {

	fmt.Println("Connecting....")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Connected!")

	return client
}