package entities

import "time"

type Note struct {
	NoteId  string `json:"note_id" gorm:"column:note_id;type:uuid;primaryKey;default:gen_random_uuid();"`
	Title   string `json:"title" gorm:"type:varchar(255);not null;"`
	Content string `json:"content" gorm:"type:text;not null;"`
  CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
  UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
