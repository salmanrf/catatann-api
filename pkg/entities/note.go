package entities

import "gorm.io/gorm"

type Note struct {
  gorm.Model
  Title string `gorm:"type:varchar(255);not null;"`
  Content string `gorm:"type:text;not null;"`
}