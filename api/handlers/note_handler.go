package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/salmanfr/catatann-api/api/presenters"
	"github.com/salmanfr/catatann-api/pkg/models"
	"github.com/salmanfr/catatann-api/pkg/note"
)

func AddNote(s note.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		dto, valid := c.Locals("dto").(models.CreateNoteDto)

		user_id := c.Locals("decoded").(jwt.MapClaims)["sub"].(string)
		
		dto.UserId = user_id
		
		if !valid {
			c.Status(http.StatusBadRequest)

			return c.JSON(presenters.NoteCustomErrorResponse(models.CreateCustomHttpError(http.StatusBadRequest, "invalid request parameters")))
		}
		
		note, err := s.InsertNote(&dto)

		if err != nil {
			c.Status(err.Code)

			return c.JSON(presenters.NoteCustomErrorResponse(err))
		}

		return c.JSON(presenters.NoteSuccessResponse(note))
	}
}

func FindOneNote(s note.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		note_id := c.Params("note_id")

		user_id := c.Locals("decoded").(jwt.MapClaims)["sub"].(string)

		note, err := s.FindOneNote(note_id, user_id)

		if err != nil {
			c.Status(err.Code)

			return c.JSON(presenters.NoteCustomErrorResponse(err))
		}

		return c.JSON(presenters.NoteSuccessResponse(note))
	}
}

func SearchNotes(s note.Service) fiber.Handler {
	return func (c *fiber.Ctx) error {
		dto, valid := c.Locals("dto").(models.SearchNoteDto)

		user_id := c.Locals("decoded").(jwt.MapClaims)["sub"].(string)
		
		if !valid {
			c.Status(http.StatusBadRequest)

			return c.JSON(presenters.NoteCustomErrorResponse(models.CreateCustomHttpError(http.StatusBadRequest, "invalid request parameters")))
		}

		dto.UserId = user_id
		
		res, err := s.SearchNotes(dto)

		if err != nil {
			c.Status(err.Code)

			return c.JSON(presenters.NoteCustomErrorResponse(err))
		}

		return c.JSON(presenters.NoteSuccessResponse(res))
	}
}

func FindNotes(s note.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		findDto, valid := c.Locals("dto").(models.FindNoteDto)

		user_id := c.Locals("decoded").(jwt.MapClaims)["sub"].(string)
		
		findDto.UserId = user_id
		
		if !valid {
			c.Status(http.StatusBadRequest)

			return c.JSON(presenters.NoteCustomErrorResponse(models.CreateCustomHttpError(http.StatusBadRequest, "invalid request parameters")))
		}
		
		result, err := s.FindNotes(findDto)

		if err != nil {
			c.Status(err.Code)

			return c.JSON(presenters.NoteCustomErrorResponse(err))
		}

		return c.JSON(presenters.NoteSuccessResponse(result))
	}
}

func UpdateNote(s note.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		note_id := c.Params("note_id");

		user_id := c.Locals("decoded").(jwt.MapClaims)["sub"].(string)
		
		updateDto, valid := c.Locals("dto").(models.UpdateNoteDto);

		updateDto.UserId = user_id
		
		if !valid {
			c.Status(http.StatusBadRequest)

			return c.JSON(presenters.NoteCustomErrorResponse(models.CreateCustomHttpError(http.StatusBadRequest, "invalid request parameters")))
		}
		
		note, err := s.UpdateNote(note_id, updateDto)

		if err != nil {
			c.Status(err.Code)

			return c.JSON(presenters.NoteCustomErrorResponse(err))
		}
		
		return c.JSON(presenters.NoteSuccessResponse(note))
	}
}

func DeleteNote(s note.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		note_id := c.Params("note_id")

		user_id := c.Locals("decoded").(jwt.MapClaims)["sub"].(string)

		note, err := s.DeleteNote(note_id, user_id)

		if err != nil {
			c.Status(err.Code)

			return c.JSON(presenters.NoteCustomErrorResponse(err))
		}

		return c.JSON(presenters.NoteSuccessResponse(note))
	}
}
