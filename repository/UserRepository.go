package repository

import "auth-golang/domain"

type UserRepository interface {
	FindAll() ([]domain.User, error)
	FindById(id int) (domain.User, error)
	Save(user domain.User) (domain.User, error)
}
