package routes

import (
	"digitalsign-api/api/controllers"
	"digitalsign-api/api/middlewares"
	"digitalsign-api/infrastructure"
)

// UserRoute -> struct
type UserRoute struct {
	logger     infrastructure.Logger
	handler    infrastructure.RequestHandler
	controller controllers.UserController
}

// Setup user routes

func (u UserRoute) Setup() {
	u.logger.Zap.Info("Setting up user routes")
	users := u.handler.Gin.Group("/user")
	{
		users.GET("", middlewares.SetMiddlewareAuthentication(), u.controller.GetAllUsers)
		users.POST("/signin", u.controller.CreateUser)
		users.POST("/login", u.controller.Login)
		users.GET("/:id", middlewares.SetMiddlewareAuthentication(), u.controller.GetUserByID)

	}
}

// NewUserRoutes -> create new user routes
func NewUserRoutes(
	logger infrastructure.Logger,
	handler infrastructure.RequestHandler,
	controller controllers.UserController,

) UserRoute {
	return UserRoute{
		logger:     logger,
		handler:    handler,
		controller: controller,
	}
}
