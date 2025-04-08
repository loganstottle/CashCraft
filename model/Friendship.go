package model

import (
	"github.com/jinzhu/gorm"
)

type Friendship struct {
	gorm.Model
	ID1    uint `json:"id1"`
	ID2    uint `json:"id2"`
	mutual uint `json:"mutual"`
}

func FriendRequest(u1 *User, id uint) {
	friend := Friendship{}
	// TODO: verify SQL algebra below
	if err := DB.First(&friend, "((id1 = ? OR id2 = ?) AND (id1 = ? OR id2 = ?))", u1.ID, u1.ID, id, id).Error; err != nil {
		// initiate friendship request
		friend := Friendship{
			ID1: u1.ID,
			ID2: id,
			mutual: 0, // false
		}
		DB.Create(&friend)
	} else if friend.ID1 != u1.ID {
		// verify friendship
		friend.mutual = 1 // true
		DB.Save(&friend)
	}
}

