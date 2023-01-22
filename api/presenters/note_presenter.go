package presenters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salmanfr/catatann-api/pkg/entities"
)

type Note struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

func NoteSuccessResponse(data *entities.Note) *fiber.Map {
	note := Note{
		Title: data.Title,
		Content: data.Content,
	}

	return &fiber.Map{
		"status": true,
		"data": note,
		"error": nil,
	}
}

func NoteErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data": nil,
		"error": err.Error(),
	}
}