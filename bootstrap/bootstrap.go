package bootstrap

import (
	"context"
	"digitalsign-api/api/controllers"
	"digitalsign-api/api/middlewares"
	"digitalsign-api/api/repository"
	"digitalsign-api/api/routes"
	"digitalsign-api/api/services"
	"digitalsign-api/infrastructure"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controllers.Module,
	middlewares.Module,
	routes.Module,
	infrastructure.Module,
	services.Module,
	repository.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler infrastructure.RequestHandler,
	routes routes.Routes,
	env infrastructure.Env,
	logger infrastructure.Logger,
	database infrastructure.Database,
	migrations infrastructure.Migrations,
) {
	conn, _ := database.DB.DB()
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("----------------------------")
			logger.Zap.Info("------- My APP V1 -------")
			logger.Zap.Info("----------------------------")
			migrations.Migrate()

			conn.SetMaxOpenConns(10)
			go func() {
				routes.Setup()

				if env.ServerPort == "" {
					handler.Gin.Run()
				} else {
					handler.Gin.Run(env.ServerPort)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Zap.Info("Stopping Application")
			conn.Close()
			return nil
		},
	})
}
