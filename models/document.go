package models

import "time"

// BaseCourse represents an item in OpenSearch
type BaseCourse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Course represents an item in DynamoDB
type Course struct {
	BaseCourse
	CreatedDate time.Time `json:"created_date"` // Creation date of the item
	Owner       string    `json:"owner"`        // Owner / Creator of the course
}
