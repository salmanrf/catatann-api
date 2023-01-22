package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salmanfr/catatann-api/api/handlers"
	"github.com/salmanfr/catatann-api/api/middlewares"
	"github.com/salmanfr/catatann-api/pkg/note"
)

func NoteRouter(app fiber.Router, s note.Service) {
	app.Get(
		"/:note_id", 
		handlers.FindOneNote(s),
	)
	app.Post(
		"/", 
		middlewares.ValidateBody(note.CreateNoteDto{}),
		handlers.AddNote(s),
	)
}