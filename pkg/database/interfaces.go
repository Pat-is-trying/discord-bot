package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type DbConnected interface {
	GetDoc() primitive.D
}