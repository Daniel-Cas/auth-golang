package main

import (
	"auth-golang/api/route"
	"auth-golang/config"
	"auth-golang/database"
)

func main() {
	cassandraSession, _ := database.Connect()
	router := route.SetupRouter()

	repository := config.ConfigRepository{Session: cassandraSession}

	repository.Init()

	defer cassandraSession.Close()

	router.Run(":8080")
}
