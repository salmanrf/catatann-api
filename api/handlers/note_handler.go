package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/salmanfr/catatann-api/api/presenters"
	"github.com/salmanfr/catatann-api/pkg/models"
	"github.com/salmanfr/catatann-api/pkg/note"
)

func AddNote(s note.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		dto, valid := c.Locals("dto").(models.CreateNoteDto)

		if !valid {
			c.Status(http.StatusBadRequest)

			return c.JSON(presenters.NoteErrorResponse(errors.New("invalid request parameters")))
		}
		
		note, err := s.InsertNote(&dto)

		if err != nil {
			c.Status(http.StatusInternalServerError)

			return c.JSON(presenters.NoteErrorResponse(err))
		}

		return c.JSON(presenters.NoteSuccessResponse(note))
	}
}

func FindOneNote(s note.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		note_id := c.Params("note_id")

		note, err := s.FindOneNote(note_id)

		if note == nil {
			c.Status(http.StatusNotFound)

			return c.JSON(presenters.NoteErrorResponse(errors.New("note not found")))
		}

		if err != nil {
			c.Status(http.StatusInternalServerError)

			return c.JSON(presenters.NoteErrorResponse(err))
		}

		return c.JSON(presenters.NoteSuccessResponse(note))
	}
}

func FindNotes(s note.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		findDto, valid := c.Locals("dto").(models.FindNoteDto)

		if !valid {
			c.Status(http.StatusBadRequest)

			return c.JSON(presenters.NoteErrorResponse(errors.New("invalid request parameters")))
		}
		
		result, err := s.FindNotes(findDto)

		if err != nil {
			c.Status(http.StatusInternalServerError)

			return c.JSON(presenters.NoteErrorResponse(err))
		}

		return c.JSON(presenters.NoteSuccessResponse(result))
	}
}

func UpdateNote(s note.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		note_id := c.Params("note_id");

		updateDto, valid := c.Locals("dto").(models.UpdateNoteDto);

		if !valid {
			c.Status(http.StatusBadRequest)

			return c.JSON(presenters.NoteErrorResponse(errors.New("invalid request parameters")))
		}
		
		note, err := s.UpdateNote(note_id, updateDto)

		if err != nil {
			c.Status(http.StatusInternalServerError)

			return c.JSON(presenters.NoteErrorResponse(err))
		}
		
		return c.JSON(presenters.NoteSuccessResponse(note))
	}
}

func DeleteNote(s note.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		note_id := c.Params("note_id")

		note, err := s.DeleteNote(note_id)

		if note == nil {
			c.Status(http.StatusNotFound)

			return c.JSON(presenters.NoteErrorResponse(errors.New("note not found")))
		}
		
		if err != nil {
			c.Status(http.StatusInternalServerError)

			return c.JSON(presenters.NoteErrorResponse(errors.New("internal server error")))
		}

		return c.JSON(presenters.NoteSuccessResponse(note))
	}
}
