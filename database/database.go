package database

import (
	"fmt"
	"url_shortener/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:mangaka2004@tcp(localhost:3306)/shortener?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=false"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	db.AutoMigrate(&models.User{}, &models.Link{})

	DB = db
}
