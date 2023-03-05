package note

import (
	"fmt"
	"net/http"

	"github.com/salmanfr/catatann-api/pkg/common"
	"github.com/salmanfr/catatann-api/pkg/entities"
	"github.com/salmanfr/catatann-api/pkg/models"
	"gorm.io/gorm"
)

type Service interface {
	InsertNote(note *models.CreateNoteDto) (*entities.Note, *models.CustomHttpErrors)
	FindOneNote(note_id string, user_id string) (*entities.Note, *models.CustomHttpErrors)
	FindNotes(dto models.FindNoteDto) (*models.Pagination, *models.CustomHttpErrors)
	UpdateNote(note_id string, dto models.UpdateNoteDto) (*entities.Note, *models.CustomHttpErrors)
	DeleteNote(note_id string, user_id string) (*entities.Note, *models.CustomHttpErrors)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		db,
	}
}

func (s *service) InsertNote(dto *models.CreateNoteDto) (*entities.Note, *models.CustomHttpErrors) {
	note := &entities.Note{
		Title: dto.Title,
		Content: dto.Content,
		UserId: dto.UserId,
	}
	
	result := s.db.Create(note)

	if result.Error != nil {
		return nil, models.CreateCustomHttpError(http.StatusInternalServerError, "internal server error")
	}

	return note, nil
}

func (s *service) FindOneNote(note_id string, user_id string) (*entities.Note, *models.CustomHttpErrors) {
	var note entities.Note

	result := s.db.First(&note, "note_id = ? AND user_id = ?", note_id, user_id)

	if result.Error != nil {
		return nil, models.CreateCustomHttpError(http.StatusNotFound, "note not found")
	}

	return &note, nil
}

func (s *service) FindNotes(dto models.FindNoteDto) (*models.Pagination, *models.CustomHttpErrors) {
	var notes []entities.Note

	pagination := &models.Pagination{
		Limit: dto.Limit,
		Page:  dto.Page,
	}

	query := s.db.Scopes(common.Paginate(notes, pagination, s.db))

	if dto.UserId != "" {
		query.Where("user_id = ?", dto.UserId)
	}
	
	if dto.Title != "" {
		query.Where("title ILIKE ?", fmt.Sprintf("%%%s%%", dto.Title))
	}

	if dto.Content != "" {
		query.Where("content ILIKE ?", fmt.Sprintf("%%%s%%", dto.Content))
	}

	res := query.Find(&notes)

	if res.Error != nil {
		return nil, models.CreateCustomHttpError(http.StatusInternalServerError, "internal server error")
	}

	pagination.Items = notes

	return pagination, nil
}

func (s *service) UpdateNote(note_id string, dto models.UpdateNoteDto) (*entities.Note, *models.CustomHttpErrors) {
	var note entities.Note

	result := s.db.First(&note, "note_id = ? AND user_id = ?", note_id, dto.UserId)

	if result.Error != nil {
		return nil, models.CreateCustomHttpError(http.StatusNotFound, "note not found")
	}

	if dto.Title != "" {
		note.Title = dto.Title
	}

	if dto.Content != "" {
		note.Content = dto.Content
	}
	
	result = s.db.Save(&note)
	
	if result.Error != nil {
		return nil, models.CreateCustomHttpError(http.StatusInternalServerError, "internal server error")
	}
	
	return &note, nil
}

func (s *service) DeleteNote(note_id string, user_id string) (*entities.Note, *models.CustomHttpErrors) {
	var note entities.Note

	result := s.db.First(&note, "note_id = ? AND user_id = ?", note_id, user_id)

	if result.Error != nil {
		return nil, models.CreateCustomHttpError(http.StatusNotFound, "note not found")
	}

	result = s.db.Delete(&note)

	if result.Error != nil {
		return nil, models.CreateCustomHttpError(http.StatusInternalServerError, "internal server error")
	}

	return &note, nil
}

