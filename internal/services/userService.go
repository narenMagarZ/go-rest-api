package services

import (
	"rest-api/internal/models"
	"rest-api/internal/repositories"
)


type UserService interface {
	Create()
	FindById(id int)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{userRepository: repository}
}

func (s *userService) Create() {
	user := models.User{}
	s.userRepository.Create(user);
}

func (s *userService) FindById(id int) {
	s.userRepository.FindById(id);
}