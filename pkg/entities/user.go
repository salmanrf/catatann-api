package entities

import "time"

type User struct {
	UserId   string `json:"user_id" gorm:"column:user_id;type:uuid;primaryKey;default:gen_random_uuid();"`
	Email    string `json:"email" gorm:"column:email;unique;not null"`
	FullName string `json:"full_name" gorm:"column:full_name;type:varchar(255);not null"`
	Password string `json:"password" gorm:"column:password;type:varchar(255);"`
	Provider string `json:"provider" gorm:"column:provider;type:varchar(255);not null;default:'local'"`
	PictureUrl string `json:"picture_url" gorm:"column:picture_url;type:varchar(255)"` 
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
  UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
  DisabledAt *time.Time `json:"disabled_at" gorm:"type:timestamp without time zone;"`
	Notes []Note `json:"notes" gorm:"foreignKey:UserId;references:UserId"`
}
