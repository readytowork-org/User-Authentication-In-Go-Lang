package infrastructure

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database model
type Database struct {
	DB *gorm.DB
}

// NewDatabse -> creates new database instant
func NewDatabse(env Env, zapLogger Logger) Database {
	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		zapLogger.Zap.Info("Url: ", psqlInfo)
		zapLogger.Zap.Panic(err)
	}
	zapLogger.Zap.Info("Database connection established")
	return Database{
		DB: db,
	}

}
