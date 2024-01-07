package main

import (
	"auth-golang/api/route"
	"auth-golang/config"
	"auth-golang/database/cassandra"
	"auth-golang/repository"
)

func main() {
	cassandraSession, _ := cassandra.Connect()
	userRepository := &repository.UserRepository{Session: cassandraSession}
	repositories := &config.Repositories{User: userRepository}
	repository := &config.ConfigRepository{Session: cassandraSession, Repositories: repositories}

	repository.Init()
	services := config.InjectServices(repository)
	router := route.SetupRouter(services)

	defer cassandraSession.Close()

	router.Run(":8080")
}
