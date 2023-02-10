package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salmanfr/catatann-api/api/handlers"
	"github.com/salmanfr/catatann-api/api/middlewares"
	"github.com/salmanfr/catatann-api/pkg/models"
	"github.com/salmanfr/catatann-api/pkg/note"
)

func NoteRouter(app fiber.Router, s note.Service) {
	app.Get(
		"/:note_id",
		handlers.FindOneNote(s),
	)
	app.Get(
		"/",
		middlewares.ValidateQuery(models.FindNoteDto{}),
		handlers.FindNotes(s),
	)
	app.Post(
		"/",
		middlewares.ValidateBody(models.CreateNoteDto{}),
		handlers.AddNote(s),
	)
	app.Put(
		"/:note_id",
		middlewares.ValidateBody(models.UpdateNoteDto{}),
		handlers.UpdateNote(s),
	)
	app.Delete(
		"/:note_id",
		handlers.DeleteNote(s),
	)
}
