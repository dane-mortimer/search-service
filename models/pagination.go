package models

type Pagination struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}
