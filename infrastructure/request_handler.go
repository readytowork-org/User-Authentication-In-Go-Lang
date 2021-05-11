package infrastructure

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RequestHandler function
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates a new request handler
func NewRequestHandler(env Env) RequestHandler {

	// Router handler
	httpRouter := gin.Default()
	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// HealthCheck godoc
	// @Summary Show the status of server.
	// @Description get the status of server.
	// @Tags root
	// @Accept */*
	// @Produce json
	// @Success 200 {object} map[string]interface{}
	// @Router / [get]
	httpRouter.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "User Authentication API Up and Running"})
	})
	return RequestHandler{Gin: httpRouter}
}
