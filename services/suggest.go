package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"search-service/opensearch"

	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

func Suggest(query string) ([]string, error) {

	// Define the search request
	req := opensearchapi.SearchRequest{
		Index: []string{"my-index"}, // Replace with your index name
		Body: strings.NewReader(fmt.Sprintf(`{
			"query": {
				"match": {
					"title": {
						"query": "%s",
						"analyzer": "standard"
					}
				}
			}
		}`, query)),
	}

	// Execute the search request
	res, err := req.Do(context.Background(), opensearch.Client)
	if err != nil {
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

	if len(titles) == 0 {
		return nil, fmt.Errorf("no titles found in the response")
	}

	return titles, nil
}
