package database

import (
	"goVanila/internal/model"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
	)
}
