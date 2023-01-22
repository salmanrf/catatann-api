package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/salmanfr/catatann-api/api/presenters"
	"github.com/salmanfr/catatann-api/pkg/entities"
	"github.com/salmanfr/catatann-api/pkg/note"
)

func AddNote(s note.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tubuh entities.Note

		err := c.BodyParser(&tubuh)

		if err != nil {
			c.Status(http.StatusBadRequest)
			
			return c.JSON(presenters.NoteErrorResponse(err))
		}
		
		note, err := s.InsertNote(&tubuh)

		if err != nil {
			c.Status(http.StatusInternalServerError)

			return c.JSON(presenters.NoteErrorResponse(err))
		}

		return c.JSON(presenters.NoteSuccessResponse(note))
	}
}

func FindOneNote(s note.Service) fiber.Handler {
	return func (c *fiber.Ctx) error {
		
		fmt.Println("GET NOTE")
		fmt.Println(c)
		
		note_id, err := strconv.Atoi(c.Params("note_id"))

		fmt.Println("note_id", note_id)
		
		if err != nil {
			c.Status(http.StatusBadRequest)

			return c.JSON(presenters.NoteErrorResponse(err))
		}

		note, err := s.GetNote(uint(note_id))

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