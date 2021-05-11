package main

import (
	"digitalsign-api/bootstrap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

// @title User Authentication API
// @version 1.0
// @description This is a authentication server.
// @termsOfService http://swagger.io/terms/

// @contact.name Ready to work
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
// @BasePath /
// @schemes http
func main() {
	godotenv.Load()
	fx.New(bootstrap.Module).Run()
}
