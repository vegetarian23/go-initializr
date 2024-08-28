package service

import (
	"github.com/vegetarian23/go-initializr/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repository.NewUserRepository(),
	}
}

func (us *UserService) GetUserById() string {
	return us.userRepository.GetUserById()
}