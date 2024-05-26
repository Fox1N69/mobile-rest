package database

import (
	"mobile/internal/app/models"

	"gorm.io/gorm"
)

func AutoMigrateScheduleModels(db *gorm.DB) error {
	tx := db.Begin()

	if err := tx.Error; err != nil {
		return err
	}

	migrateModels := []interface{}{
		&models.Prepod{},
		&models.Group{},
		&models.Subject{},
		&models.Change{},
		&models.Urok{},
	}

	for _, model := range migrateModels {
		if err := tx.AutoMigrate(model); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
