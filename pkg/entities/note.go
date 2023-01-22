package entities

import "gorm.io/gorm"

type Note struct {
  gorm.Model
	NoteId string `json:"note_id" gorm:"field:note_id;type:uuid;primaryKey;default:gen_random_uuid();"`
  Title string `gorm:"type:varchar(255);not null;"`
  Content string `gorm:"type:text;not null;"`
}