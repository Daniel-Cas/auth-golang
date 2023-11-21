package config

import (
	"auth-golang/database"
	"auth-golang/repository"
	"context"
	"github.com/gocql/gocql"
)

type ConfigRepository struct {
	Session      *gocql.Session
	Repositories *Repositories
}

type Repositories struct {
	User *repository.UserRepository
}

func (config *ConfigRepository) Init() {
	config.Repositories = InitRepositories(config.Session)

	ExecuteMigrations(config.Session)
}

func ExecuteMigrations(session *gocql.Session) {
	database.CassandraMigrations(
		context.Background(), session,
	)
}

func InitRepositories(session *gocql.Session) *Repositories {
	return &Repositories{
		User: &repository.UserRepository{
			Session: session,
		},
	}
}
