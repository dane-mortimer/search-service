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
	req := opensearchapi.SearchRequest{
		Body: strings.NewReader(fmt.Sprintf(`{
					"suggest": {
							"my-suggestion": {
									"text": "%s",
									"term": {
											"field": "title"
									}
							}
					}
			}`, query)),
	}

	res, err := req.Do(context.Background(), opensearch.Client)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Log the raw response for debugging
	body, _ := io.ReadAll(res.Body)
	log.Printf("OpenSearch Response: %s", string(body))

	var result map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(string(body))).Decode(&result); err != nil {
		return nil, err
	}

	// Extract suggestions from the response
	suggest, ok := result["suggest"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid suggest response format")
	}

	mySuggestion, ok := suggest["my-suggestion"].([]interface{})
	if !ok || len(mySuggestion) == 0 {
		return nil, fmt.Errorf("no suggestions found")
	}

	// Get the options from the first suggestion
	options, ok := mySuggestion[0].(map[string]interface{})["options"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid options format")
	}

	// Extract the text from each option
	var output []string
	for _, option := range options {
		opt, ok := option.(map[string]interface{})
		if !ok {
			continue
		}
		text, ok := opt["text"].(string)
		if !ok {
			continue
		}
		output = append(output, text)
	}

	return output, nil
}
