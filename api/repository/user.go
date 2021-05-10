package repository

import (
	"digitalsign-api/infrastructure"
	"digitalsign-api/models"
	"digitalsign-api/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// UserRepository -> struct
type UserRepository struct {
	db infrastructure.Database
}

// NewUserRepository --> constructor
func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

// GetAllUsers -> func
func (u UserRepository) GetAllUsers(pagination utils.Pagination, searchParams models.UserSearchParams) ([]models.User, int64, error) {
	var users []models.User
	var count int64
	querybuilder := u.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset)
	if pagination.All {
		querybuilder = u.db.DB
	}
	if searchParams.Keyword != "" {
		query := "%" + searchParams.Keyword + "%"
		querybuilder = querybuilder.Where(
			u.db.DB.Where("username LIKE ?", query).
				Or("first_name LIKE ?", query).Or("last_name LIKE ?", query))
	}
	err := querybuilder.Model(&models.User{}).
		Order("updated_at ").
		Where(&users).
		Find(&users).
		Offset(-1).
		Limit(-1).
		Count(&count).Error
	return users, count, err
}

//CreateUser --> func
func (u UserRepository) CreateUser(user models.User) (models.User, error) {
	return user, u.db.DB.Create(&user).Error
}

// Login User ---> func
func (u UserRepository) Login(email, password string) (models.User, error) {
	var user models.User

	if resp := u.db.DB.First(&user, "email = ?", email).Error; resp != nil {
		notFound := errors.New("email not found")
		return user, notFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		invalidPassword := errors.New("invalid password")
		return user, invalidPassword
	}
	return user, nil

}
