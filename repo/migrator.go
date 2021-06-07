package repo

import (
	"fmt"
	"github.com/dipesh23-apt/golang_api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitialMigration() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open("people.db"), &gorm.Config{})
	DB = db
	if err != nil {
		fmt.Println("Failed to connect to Database")
	}
	// Drop table if exists (will ignore or delete foreign key constraints when dropping)
	x := db.Migrator().HasTable(&models.User{})
	if !x {
		db.AutoMigrate(&models.User{})
	}
	return db, nil
}
