package bootstrap

import (
	"context"
	"fx-modules/infrastructure"

	"go.uber.org/fx"
)

var Module = fx.Options(
	infrastructure.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	logger infrastructure.Logger,
	env infrastructure.Env,
	database infrastructure.Database,
	migrations infrastructure.Migrations,
) {
	conn, _ := database.DB.DB()
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("-------------------------")

			migrations.Migrate()
			conn.SetMaxOpenConns(10)
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Zap.Info("Stopping Application")
			return nil
		},
	})
}
