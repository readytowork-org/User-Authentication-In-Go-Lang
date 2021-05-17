package infrastructure

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrations -> Migration stuct
type Migrations struct {
	logger Logger
	env    Env
}

//NewMigrations -> returns new Migrations struct
func NewMigrations(logger Logger, env Env) Migrations {
	return Migrations{
		logger: logger,
		env:    env,
	}
}

// Migrate all table
func (m Migrations) Migrate() {
	m.logger.Zap.Info("------MIGRATING SCHEMAS-------")
	username := m.env.DBUsername
	password := m.env.DBPassword
	host := m.env.DBHost
	// port := m.env.DBPort
	dbname := m.env.DBName

	psqlInfo := fmt.Sprintf("%s:%s@%s/%s?sslmode=disable", username, password, host, dbname)
	migrations, err := migrate.New("file://migration/", "postgres://"+psqlInfo)
	if err != nil {
		m.logger.Zap.Error("Error migrating files:::", err.Error())
		panic(err)
	}
	if err := migrations.Up(); err != nil {
		m.logger.Zap.Error("No files found for migration:::", err.Error())
		log.Fatal(err)
	}
}
