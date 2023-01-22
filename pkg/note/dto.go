package note

type CreateNoteDto struct {
	Title string `json:"title" validate:"required,min=3,max=255"`
	Content string `json:"content" validate:"required,min=3"`
}