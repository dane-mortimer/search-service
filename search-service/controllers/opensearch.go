package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"search-service/clients"
	"search-service/models"

	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

func SearchCourseController(query, pageStr, sizeStr string, fields []string) ([]models.BaseCourse, int, error) {

	index := os.Getenv("COURSE_INDEX")

	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)

	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}

	from := (page - 1) * size

	// Construct the fields array for the multi_match query
	fieldsJSON, err := json.Marshal(fields)
	if err != nil {
		log.Printf("failed to marshal fields: %v", err)
		return nil, 0, fmt.Errorf("failed to marshal fields: %w", err)
	}

	req := opensearchapi.SearchRequest{
		Index: []string{index},
		Body: strings.NewReader(fmt.Sprintf(`{
			"query": {
				"multi_match": {
					"query": "%s",
					"type": "phrase_prefix",
					"fields": %s,
					"slop": 2
				}
			},
			"from": %d,
			"size": %d
		}`, query, string(fieldsJSON), from, size)),
	}

	res, err := req.Do(context.Background(), clients.Client)
	if err != nil {
		log.Printf("Opensearch search request failed: %v", err)
		return nil, 0, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	totalItems := int(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	var documents []models.BaseCourse
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"]
		doc := models.BaseCourse{
			ID:      hit.(map[string]interface{})["_id"].(string),
			Title:   source.(map[string]interface{})["title"].(string),
			Content: source.(map[string]interface{})["content"].(string),
		}
		documents = append(documents, doc)
	}

	return documents, totalItems, nil
}

func SuggestCourseController(query string) ([]string, error) {

	index := os.Getenv("COURSE_INDEX")
	// Define the search request
	req := opensearchapi.SearchRequest{
		Index: []string{index},
		Body: strings.NewReader(fmt.Sprintf(`{
			"query": {
				"multi_match": {
					"query": "%s",
					"type": "phrase_prefix",
					"slop": 2,
					"fields": [
						"title"
					]
				}
			},
			"size": 3
		}`, query)),
	}

	// Execute the search request
	res, err := req.Do(context.Background(), clients.Client)
	if err != nil {
		log.Printf("Error executing search request: %v", err)
		return nil, fmt.Errorf("error executing search request: %w", err)
	}
	defer res.Body.Close()

	// Log the raw response for debugging
	body, _ := io.ReadAll(res.Body)
	log.Printf("OpenSearch Response: %s", string(body))

	// Parse the response
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	// Extract hits from the response
	hits, ok := result["hits"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid hits format in response")
	}

	// Extract the list of documents
	docs, ok := hits["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid hits.hits format in response")
	}

	// Extract the titles from the documents
	var titles []string
	for _, doc := range docs {
		docMap, ok := doc.(map[string]interface{})
		if !ok {
			continue
		}

		source, ok := docMap["_source"].(map[string]interface{})
		if !ok {
			continue
		}

		title, ok := source["title"].(string)
		if !ok {
			continue
		}

		titles = append(titles, title)
	}

	return titles, nil
}
