package repository

import (
	"auth-golang/domain"
	"github.com/gocql/gocql"
)

type UserRepository struct {
	Session *gocql.Session
}

func (repository *UserRepository) Save(user *domain.User) error {
	query := "INSERT INTO users (id, username, password, created_at) VALUES (?, ?, ?, ?)"

	err := repository.Session.Query(
		query,
		user.Id,
		user.Username,
		user.Password,
		user.CreatedAt,
	).Exec()

	if err != nil {
		return err
	}

	return nil
}
