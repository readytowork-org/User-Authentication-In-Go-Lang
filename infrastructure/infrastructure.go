package infrastructure

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewEnv),
	fx.Provide(NewRequestHandler),
	fx.Provide(NewMigrations),
	fx.Provide(NewFBApp),
	fx.Provide(NewFBAuth),
	fx.Provide(NewFCMClient),
	fx.Provide(NewLogger),
	fx.Provide(NewDatabase),
	fx.Provide(NewGmailService),
)
