package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"search-service/clients"
	"search-service/models"

	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

func Search(query, pageStr, sizeStr string, fields []string) ([]models.Document, int, error) {
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
		return nil, 0, fmt.Errorf("failed to marshal fields: %w", err)
	}

	req := opensearchapi.SearchRequest{
		Body: strings.NewReader(fmt.Sprintf(`{
			"query": {
				"multi_match": {
					"query": "%s",
					"type": "bool_prefix",
					"fields": %s
				}
			},
			"from": %d,
			"size": %d
		}`, query, string(fieldsJSON), from, size)),
	}

	res, err := req.Do(context.Background(), clients.Client)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	totalItems := int(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	var documents []models.Document
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"]
		doc := models.Document{
			ID:      hit.(map[string]interface{})["_id"].(string),
			Title:   source.(map[string]interface{})["title"].(string),
			Content: source.(map[string]interface{})["content"].(string),
		}
		documents = append(documents, doc)
	}

	return documents, totalItems, nil
}
