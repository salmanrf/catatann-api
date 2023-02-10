package note

import (
	"fmt"
	"log"

	"github.com/salmanfr/catatann-api/pkg/common"
	"github.com/salmanfr/catatann-api/pkg/entities"
	"github.com/salmanfr/catatann-api/pkg/models"
	"gorm.io/gorm"
)

type Service interface {
	InsertNote(note *models.CreateNoteDto) (*entities.Note, error)
	FindOneNote(note_id string) (*entities.Note, error)
	FindNotes(dto models.FindNoteDto) (*models.Pagination, error)
	UpdateNote(note_id string, dto models.UpdateNoteDto) (*entities.Note, error)
	DeleteNote(note_id string) (*entities.Note, error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		db,
	}
}

func (s *service) InsertNote(dto *models.CreateNoteDto) (*entities.Note, error) {
	note := &entities.Note{
		Title: dto.Title,
		Content: dto.Content,
	}
	
	result := s.db.Create(note)

	if result.Error != nil {
		return nil, result.Error
	}

	return note, nil
}

func (s *service) FindOneNote(note_id string) (*entities.Note, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("ERROR at note service's FindOneNote", err)
		}
	}()

	var note entities.Note

	result := s.db.First(&note, "note_id = ?", note_id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &note, nil
}

func (s *service) FindNotes(dto models.FindNoteDto) (*models.Pagination, error) {
	var notes []entities.Note

	pagination := &models.Pagination{
		Limit: dto.Limit,
		Page:  dto.Page,
	}

	query := s.db.Scopes(common.Paginate(notes, pagination, s.db))

	if dto.Title != "" {
		query = query.Where("title ILIKE ?", fmt.Sprintf("%%%s%%", dto.Title))
	}

	if dto.Content != "" {
		query = query.Where("content ILIKE ?", fmt.Sprintf("%%%s%%", dto.Content))
	}

	query.Find(&notes)

	pagination.Items = notes

	return pagination, nil
}

func (s *service) UpdateNote(note_id string, dto models.UpdateNoteDto) (*entities.Note, error) {
	var note entities.Note

	result := s.db.First(&note, "note_id = ?", note_id)

	if result.Error != nil {
		return nil, result.Error
	}

	if dto.Title != "" {
		note.Title = dto.Title
	}

	if dto.Content != "" {
		note.Content = dto.Content
	}
	
	result = s.db.Save(&note)
	
	if result.Error != nil {
		return nil, result.Error
	}
	
	return &note, nil
}

func (s *service) DeleteNote(note_id string) (*entities.Note, error) {
	var note entities.Note

	result := s.db.First(&note, "note_id = ?", note_id)

	if result.Error != nil {
		return nil, result.Error
	}

	result = s.db.Delete(&note)

	if result.Error != nil {
		return nil, result.Error
	}

	return &note, nil
}

