package model

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
    "os"
    "fmt"
	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

func ConnectionDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Couldn't open .env")
	}
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
		fmt.Printf("hash function failed: %s", err)
	}
	return string(hashedPassword)
}
