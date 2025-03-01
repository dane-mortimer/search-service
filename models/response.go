package models

type SuccessResponse struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"` // Optional pagination data
}

// ErrorResponse represents an error JSON response
type ErrorResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code,omitempty"`
	Details interface{} `json:"details,omitempty"`
}
