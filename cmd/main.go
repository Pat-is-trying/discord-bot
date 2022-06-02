package main

import (
	"discord-bot/pkg/api/router"
	"discord-bot/pkg/session"
)

func main() {

	s := session.InitSession()
	router.SetRouter(s)
	session.Connect(s)
}