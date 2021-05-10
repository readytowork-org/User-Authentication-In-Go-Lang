package services

import (
	"digitalsign-api/api/repository"
	"digitalsign-api/models"
	"digitalsign-api/utils"
)

type UserService struct {
	repository repository.UserRepository
}

// NewUserService -> Constructor
func NewUserService(repository repository.UserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

// GetAllUsers -> return all users
func (u UserService) GetAllUsers(pagination utils.Pagination, searchParams models.UserSearchParams) ([]models.User, int64, error) {
	return u.repository.GetAllUsers(pagination, searchParams)
}

// CreateUser ---> create a new user
func (u UserService) CreateUser(user models.User) (models.User, error) {
	return u.repository.CreateUser(user)
}

//LoginUser ----> login user
func (u UserService) Login(email string, password string) (models.User, error) {
	return u.repository.Login(email, password)
}
