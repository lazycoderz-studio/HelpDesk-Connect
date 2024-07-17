package db

import (
	"fmt"
	"log"
	"time"

	"lazycoderz-studio/helpdesk-connect/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

// InitDB initializes the database connection using GORM
func InitDB(cfg *config.DatabaseConfig) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return fmt.Errorf("failed to initialize GORM: %w", err)
	}

	// Configure the connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from GORM: %w", err)
	}

	// Set the maximum number of open connections
	sqlDB.SetMaxOpenConns(50)

	// Set the maximum number of idle connections
	sqlDB.SetMaxIdleConns(10)

	// Set the maximum lifetime of a connection
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("failed to get sql.DB from GORM: %v", err)
		return
	}
	sqlDB.Close()
}

func GetDB() *gorm.DB {
	return DB
}
