package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Age         int32
	Description string
}
