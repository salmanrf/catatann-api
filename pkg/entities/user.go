package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId   string `gorm:"column:user_id;primaryKey;type:uuid;"`
	Email    string `gorm:"column:email;unique;"`
	FullName string `gorm:"column:full_name;type:varchar(255);"`
	Password string `gorm:"column:password;type:varchar(255);"`
}
