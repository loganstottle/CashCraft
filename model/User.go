package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID           int     `json:"userid"`
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	Cash         float64 `json:"cash"`
	Stocks       []Stock `json:"stocks"`
	SessionToken string  `json:"session_token"`
}
