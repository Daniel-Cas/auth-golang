package service

import (
	"auth-golang/domain"
	"auth-golang/repository"
	"github.com/gocql/gocql"
)

type UserService struct {
	Repository *repository.UserRepository
}

func (service *UserService) FindAll() ([]domain.User, error) {
	return service.Repository.FindAll()
}

func (service *UserService) FindById(id gocql.UUID) (domain.User, error) {
	return service.Repository.FindById(id)
}

func (service *UserService) Save(user *domain.User) error {
	return service.Repository.Save(user)
}
