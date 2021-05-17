package controllers

import (
	"fx-modules/infrastructure"
)

// UserController struct
type UserController struct {
	logger infrastructure.Logger
}

// NewUserController -> constructor
func NewUserController(logger infrastructure.Logger) UserController {
	return UserController{
		logger: logger,
	}
}
