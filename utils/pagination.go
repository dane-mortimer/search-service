package utils

import (
	"math"
	"strconv"
)

// Pagination represents pagination details
type Pagination struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

// NewPagination creates a new Pagination instance
func NewPagination(pageStr, sizeStr string, totalItems int) Pagination {
	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)

	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(size)))

	return Pagination{
		Page:       page,
		Size:       size,
		TotalPages: totalPages,
		TotalItems: totalItems,
	}
}
