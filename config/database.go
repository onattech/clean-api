package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgresDatabase establishes a connection to PostgreSQL using GORM.
func NewPostgresDatabase(env *Env) *gorm.DB {
	// Build the PostgreSQL connection string
	postgresURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		env.DBHost, env.DBPort, env.DBUser, env.DBPass, env.DBName, env.DBSSLMode)

	// Configure GORM with some options, including a logger
	db, err := gorm.Open(postgres.Open(postgresURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	// Optional: Set the connection pool settings (e.g., max connections, idle time)
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database object from GORM: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Ping the database to make sure the connection works
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping PostgreSQL database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL")

	// Call the migration function
	if err := MigrateDatabase(db); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	return db
}

// ClosePostgresDBConnection closes the connection to the PostgreSQL database.
func ClosePostgresDBConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database object from GORM: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("Failed to close the PostgreSQL connection: %v", err)
	}

	log.Println("Connection to PostgreSQL closed.")
}
