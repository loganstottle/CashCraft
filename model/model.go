package model

import (
	"encoding/hex"
	"fmt"
	"os"

	"crypto/sha256"

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

	database.AutoMigrate(&User{}, &Stock{}, &StockPrice{})

	DB = database
}

func HashPassword(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	return string(hashedPassword)
}
