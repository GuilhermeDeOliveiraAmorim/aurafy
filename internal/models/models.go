package models

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB, sqlDB *sql.DB) {
	if err := db.AutoMigrate(); err != nil {
		fmt.Println("Error migrating database:", err)
		return
	}
}
