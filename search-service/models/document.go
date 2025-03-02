package models

// BaseCourse represents an item in OpenSearch
type BaseCourse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Course struct {
	ID        string `dynamodbav:"id"`
	Title     string `dynamodbav:"title"`
	Content   string `dynamodbav:"content"`
	CreatedAt string `dynamodbav:"created_at"` // Creation date of the item
	Owner     string `dynamodbav:"owner"`      // Owner / Creator of the course
}
