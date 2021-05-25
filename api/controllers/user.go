package controllers

import (
	"digitalsign-api/api/responses"
	"digitalsign-api/api/services"
	"digitalsign-api/infrastructure"
	"digitalsign-api/models"
	"digitalsign-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserController -> struct
type UserController struct {
	logger   infrastructure.Logger
	services services.UserService
}

//NewUserController -> constructor
func NewUserController(
	logger infrastructure.Logger,
	us services.UserService,
) UserController {
	return UserController{
		logger:   logger,
		services: us,
	}
}

// GetAllUsers -> get all users
func (u UserController) GetAllUsers(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	searchParams := models.UserSearchParams{
		Keyword: c.Query("keyword"),
	}
	users, count, err := u.services.GetAllUsers(pagination, searchParams)

	if err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to get users")
		return
	}
	var user_list []interface{}
	// To eliminate the passsword field
	for _, e := range users {
		user := e.ToMap()
		user_list = append(user_list, user)

	}
	responses.JSONCount(c, http.StatusOK, user_list, int(count))

}

// CreateUser ----> creates a new user
func (u UserController) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		u.logger.Zap.Error("User body parse error in controller:", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to Create User")
		return
	}

	pagination := utils.BuildPagination(c)
	searchParams := models.UserSearchParams{
		Keyword: c.Query("keyword"),
	}
	// test for duplicate email and username
	all_users, _, _ := u.services.GetAllUsers(pagination, searchParams)
	for _, e := range all_users {
		if e.Email == user.Email {
			responses.ErrorJSON(c, http.StatusBadRequest, "Email is taken")
			return
		}
		if e.Username == user.Username {
			responses.ErrorJSON(c, http.StatusBadRequest, "Username is taken")
			return
		}
	}
	// Hash and salt the password
	if user.Password != "" {
		user.Password, _ = utils.HashAndSalt([]byte(user.Password))
	}
	err := user.Validate("update")
	if err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}
	if _, err := u.services.CreateUser(user); err != nil {
		u.logger.Zap.Error("Failed to save user in service:", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to Save User")
		return
	}
	responses.SuccessJSON(c, http.StatusCreated, "User Added Successfully")
}

// Login user --->
func (u UserController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		u.logger.Zap.Error("User body parse error in controller:", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, err.Error())
	}
	validation_err := user.Validate("login")
	if validation_err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, validation_err.Error())
		return
	}
	response, service_err := u.services.Login(user.Email, user.Password)
	if service_err != nil {
		u.logger.Zap.Error("Failed to get user:", service_err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, service_err.Error())
		return
	}

	jwt_token, err := utils.CreateToken(uint32(user.ID))

	if err != nil {
		u.logger.Zap.Error("Error creating Token", err.Error())
		responses.ErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}
	// Called to map method to remove password field
	result := response.ToMap()
	c.JSON(http.StatusOK, gin.H{
		"user":  result,
		"token": jwt_token,
	})
}

// GetUserByID
func (u UserController) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to parse id")
		return
	}
	user, err := u.services.GetUserByID(id)
	if err != nil {
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to get user with given ID")
		return
	}
	responses.ErrorJSON(c, http.StatusBadRequest, user)

}
