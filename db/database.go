package db

import (
	"fmt"

	"github.com/go-mpesa-integration/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvCofig , DBMigrator func(db *gorm.DB) error)*gorm.DB {
	databaseURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode)

	//connect to database
	
	db , err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Fatalf("Failed to connect to database: %s", err)
	}
	logrus.Info("Database connection established successfully 🚀")

	//run migrations

	if err := DBMigrator(db); err != nil {
		logrus.Fatalf("Failed to run migrations: %s", err)
	}
	logrus.Info("Migrations run successfully 🚀")
	return db 
}