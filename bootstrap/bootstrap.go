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
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("-------------------------")
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Zap.Info("Stopping Application")
			return nil
		},
	})
}
