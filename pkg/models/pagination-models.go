package models

type PaginationRequest struct {
	Limit      int         `json:"limit" query:"limit" validate:"numeric"`
	Page       int         `json:"page" query:"page" validate:"numeric,gte=1"`
	SortField  string      `json:"sort_field" query:"sort_field"`
	SortOrder  string      `json:"sort_order" query:"sort_order"`
}

type Pagination struct {
	Limit      int         `json:"limit" query:"limit"`
	Page       int         `json:"page" query:"page"`
	SortField  string      `json:"sort_field" query:"sort_field"`
	SortOrder  string      `json:"sort_order" query:"sort_order"`
	TotalItems int64       `json:"total_items" query:"sort_field"`
	TotalPages int         `json:"total_pages"`
	Items      interface{} `json:"items"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}

	return p.Page
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}

	return p.Limit
}

func (p *Pagination) GetSortField() string {
	if p.SortField == "" {
		p.SortField = "created_at"
	}

	return p.SortField
}

func (p *Pagination) GetSortOrder() string {
	if p.SortOrder == "" {
		p.SortOrder = "DESC"
	}

	return p.SortOrder
}