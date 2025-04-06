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
	user := os.Getenv("DBUSER") // Getting data out of the .env file - same next line
	password := os.Getenv("DBPW")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/cashcraft?parseTime=true", user, password) // Configures the connection to the database
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})                              // connection to the database

	if err != nil {
		panic("Failed to connect to database!") // Couldnt be more clearn than this, if error isnt no, go AGGGGHHHHHHH
	}

	database.AutoMigrate(&User{}, &Stock{}, &StockPrice{}) // This is something that automatically turns structs into database - very cool

	DB = database // I question this line, but I do not touch it (Care to explain this anyone? - The Ginger)
}

func HashPassword(input string) string { // sha256 hashing for passwords - the industry standard
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	return string(hashedPassword)
}
