package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

//UserSearch Params -> search user params
type UserSearchParams struct {
	Keyword string
}

// User -> model
type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Base
}

// TableName -> TableName
func (u User) TableName() string {
	return "user"
}

// ToMap -> convert to map
func (u User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         u.ID,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"username":   u.Username,
		"email":      u.Email,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	}
}

// Validate user data
func (u *User) Validate(action string) error {
	var err error
	switch strings.ToLower(action) {
	case "update":
		if u.Username == "" {
			err = errors.New("required Nickname")
			return err
		}
		if u.Password == "" {
			err = errors.New("required Password")
			return err
		}
		if u.Email == "" {
			err = errors.New("required Email")
			return err
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			err = errors.New("invalid Email")
			return err
		}

		return err
	case "login":
		if u.Password == "" {
			err = errors.New("required Password")
			return err
		}
		if u.Email == "" {
			err = errors.New("required Email")
			return err
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			err = errors.New("invalid Email")
			return err
		}
		return err

	default:
		if u.Username == "" {
			err = errors.New("required Nickname")
			return err
		}
		if u.Password == "" {
			err = errors.New("required Password")
			return err
		}
		if u.Email == "" {
			err = errors.New("required Email")
			return err
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			err = errors.New("invalid Email")
			return err
		}
		return nil
	}
}
