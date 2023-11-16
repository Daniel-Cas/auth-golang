package service

import (
	"auth-golang/domain"
	"auth-golang/repository"
)

type UserService struct {
	Repository repository.UserRepository
}

func (service *UserService) FindAll() ([]domain.User, error) {
	return service.Repository.FindAll()
}

func (service *UserService) FindById(id int) (domain.User, error) {
	return service.Repository.FindById(id)
}

func (service *UserService) Save(user domain.User) (domain.User, error) {
	return service.Repository.Save(user)
}
