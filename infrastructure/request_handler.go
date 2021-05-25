package infrastructure

import (
	"fmt"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"

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
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://dedec96d410545aca5889f9a8b083f85@o722130.ingest.sentry.io/5781600",
	}); err != nil {
		fmt.Println("Sentry initialization failed....")
	}
	httpRouter.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("Sentry is initialized!")

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
