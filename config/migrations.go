package config

import (
	"errors"
	"fmt"
	"log"

	"github.com/onattech/invest/models"
	"gorm.io/gorm"
)

type SchemaVersion struct {
	ID      int `gorm:"primaryKey"`
	Version int
}

const currentSchemaVersion = 1

func MigrateDatabase(db *gorm.DB) error {
	// Ensure the SchemaVersion table exists
	if !db.Migrator().HasTable(&SchemaVersion{}) {
		if err := db.AutoMigrate(&SchemaVersion{}); err != nil {
			return fmt.Errorf("failed to migrate SchemaVersion: %w", err)
		}
	}

	// Check and update schema version
	var schemaVersion SchemaVersion
	err := db.First(&schemaVersion, "id = ?", 1).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		schemaVersion = SchemaVersion{ID: 1, Version: 0}
		if err := db.Create(&schemaVersion).Error; err != nil {
			return fmt.Errorf("failed to create schema version record: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to retrieve schema version: %w", err)
	}

	if schemaVersion.Version < currentSchemaVersion {
		log.Println("Initiating migration...")
		err := db.Transaction(func(tx *gorm.DB) error {
			// Acquire advisory lock to prevent concurrent migrations
			if err := db.Exec("SELECT pg_advisory_lock(1)").Error; err != nil {
				return fmt.Errorf("failed to acquire advisory lock: %w", err)
			}
			defer func() {
				if err := db.Exec("SELECT pg_advisory_unlock(1)").Error; err != nil {
					log.Printf("failed to release advisory lock: %v", err)
				}
			}()

			// Perform migrations
			if err := tx.AutoMigrate(&models.User{} /*, Add other models here */); err != nil {
				return fmt.Errorf("failed to migrate models: %w", err)
			}

			// Update schema version
			if err := tx.Model(&schemaVersion).Update("Version", currentSchemaVersion).Error; err != nil {
				return fmt.Errorf("failed to update schema version to %d: %w", currentSchemaVersion, err)
			}
			return nil
		})
		if err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
		log.Println("Migration Completed...")
	} else {
		log.Println("Database schema is up to date.")
	}

	return nil
}
