package models

type CreateNoteDto struct {
	Title   string `json:"title" validate:"required,min=3,max=255"`
	Content string `json:"content" validate:"required,min=3"`
	UserId string `json:"user_id"`
}

type UpdateNoteDto struct {
	Title string `json:"title" validate:"min=3,max=255"`
	Content string `json:"content" validate:"min=3"`
	UserId string `json:"user_id"`
}

type FindNoteDto struct {
	PaginationRequest
	Title   string `json:"title" query:"title"`
	Content string `json:"content" query:"content"`
	UserId  string `json:"user_id" query:"user_id"`
}
