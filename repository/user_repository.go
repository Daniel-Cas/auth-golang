package repository

import (
	"auth-golang/domain"
	"github.com/gocql/gocql"
	"time"
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

func (repository *UserRepository) FindAll() ([]domain.User, error) {
	var users []domain.User

	query := "SELECT id, username, password, created_at FROM users"
	iterable := repository.Session.Query(query).Consistency(gocql.All).Iter()

	var id string
	var username string
	var password string
	var createdAt time.Time

	for iterable.Scan(&id, &username, &password, &createdAt) {
		users = append(users, domain.User{Id: id, Username: username, Password: password, CreatedAt: createdAt})
	}

	if err := iterable.Close(); err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *UserRepository) FindById(id gocql.UUID) (domain.User, error) {
	var user domain.User

	query := "SELECT id, username, password, created_at FROM users WHERE id = ?"

	var userId string
	var username string
	var password string
	var createdAt time.Time

	err := repository.Session.Query(query, id).Consistency(gocql.One).Scan(&userId, &username, &password, &createdAt)
	{
		user = domain.User{Id: userId, Username: username, Password: password, CreatedAt: createdAt}
	}

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
