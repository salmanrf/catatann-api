package entities

import "time"

type Token struct {
	TokenId string `json:"token_id" gorm:"column:token_id;type:uuid;primaryKey;default:gen_random_uuid();"`
	UserId string `json:"user_id" gorm:"column:user_id;type:uuid;not null"`
	ExpiredAt *time.Time `json:"expired_at" gorm:"column:expired_at;type:timestamp without time zone;"`
}
