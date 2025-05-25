package services

import (
	"rest-api/internal/models"
	"rest-api/internal/repositories"
	"rest-api/internal/types"
)

type UserService interface {
	Create(user models.User) error
	FindById(id int) (*models.User, error)
	FindOne(condition models.User) (*models.User, error)
	DeleteOne(condition struct{ Id int }) error
	UpdateOne(id int, input models.User) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{userRepository: repository}
}

func (s *userService) Create(user models.User) error {
	return s.userRepository.Create(user)
}

func (s *userService) FindById(id int) (*models.User, error) {
	return s.userRepository.FindById(id)
}

func (s *userService) FindOne(condition models.User) (*models.User, error) {
	return s.userRepository.FindOne(condition)
}

func (s *userService) FindAll(condition models.User) ([]*models.User, error) {
	return s.userRepository.FindAll(types.CursorPaginationArgs{})
}

func (s *userService) CursorPagination() {
	where := map[string]interface{}{
		"email":    "",
		"username": "",
	}
	go s.userRepository.Count(where)
	go s.userRepository.FindAll(types.CursorPaginationArgs{
		Where: where,
		Limit: 10,
		Order: "id",
		Sort:  true,
	})

}

func (s *userService) DeleteOne(condition struct{ Id int }) error {
	return s.userRepository.Delete(condition.Id)
}

func (s *userService) UpdateOne(id int, input models.User) error {
	return s.userRepository.UpdateOne(id, input)
}
