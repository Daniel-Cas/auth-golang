package config

import (
	"auth-golang/service"
)

type ConfigService struct {
	Services *Services
}

type Services struct {
	User *service.UserService
}

func InjectServices(configRepository *ConfigRepository) *Services {
	return &Services{
		User: &service.UserService{
			Repository: configRepository.Repositories.User,
		},
	}
}
