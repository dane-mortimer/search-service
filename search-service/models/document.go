package models

// BaseCourse represents an item in OpenSearch
type BaseCourse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Course struct {
	ID        string `dynamodbav:"id" json:"id" validate:"omitempty"`
	Title     string `dynamodbav:"title" json:"title" validate:"required,min=3,max=100"`
	Content   string `dynamodbav:"content" json:"content" validate:"required,min=10"`
	CreatedAt string `dynamodbav:"created_at" json:"created_at" validate:"omitempty"`
	Owner     string `dynamodbav:"owner" json:"owner" validate:"required,min=3,max=50"`
}
