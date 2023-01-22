package note

import (
	"fmt"
	"log"

	"github.com/salmanfr/catatann-api/pkg/entities"
	"gorm.io/gorm"
)

type Service interface {
	InsertNote(note *entities.Note) (*entities.Note, error)
	GetNote(id uint) (*entities.Note, error)
}

type service struct{
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		db,
	}
}

func (s *service) InsertNote(note *entities.Note) (*entities.Note, error) {
	result := s.db.Create(note)

	if result.Error != nil {
		return nil, result.Error
	}
	
	return note, nil
}


func (s *service) GetNote(id uint) (*entities.Note, error) {
	defer func () {
		if err := recover(); err != nil {
			log.Println("ERROR at note service's GetNote", err)
		}
	}()
	
	var note entities.Note

	fmt.Println("note ", note)
	fmt.Println("db ", s.db)
	
	result := s.db.First(&note, id)

	if result.Error != nil {
		return nil, result.Error
	}
	
	return &note, nil
}

