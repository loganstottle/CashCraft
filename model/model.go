package model

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPW")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/cashcraft?parseTime=true", user, password)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&User{})

	DB = database
}

func HashPassword(input string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("hash function failed: %s", err)
		return ""
	}
	return string(hashedPassword)
}
