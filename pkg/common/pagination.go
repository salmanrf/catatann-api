package common

import (
	"fmt"
	"math"
	"strings"

	"github.com/salmanfr/catatann-api/pkg/models"
	"gorm.io/gorm"
)

func Paginate(value interface{}, pagination *models.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalItems int64

	if pagination.Limit == 0 {
		pagination.Limit = 10
	}
	
	db.Model(value).Count(&totalItems)

	pagination.TotalItems = totalItems
	totalPages := int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))

	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		sort := fmt.Sprintf("%s %s", pagination.GetSortField(), strings.ToUpper(
			pagination.GetSortOrder(),
		))

		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(sort)
	}
}
