package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	User         string `json:"user"`
	Password     string `json:"password"`
	SessionToken string `json:"session_token"`
}
