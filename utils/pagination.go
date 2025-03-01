package utils

import (
	"math"
	"search-service/models"
	"strconv"
)

// NewPagination creates a new Pagination instance
func NewPagination(pageStr, sizeStr string, totalItems int) models.Pagination {
	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)

	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(size)))

	return models.Pagination{
		Page:       page,
		Size:       size,
		TotalPages: totalPages,
		TotalItems: totalItems,
	}
}
