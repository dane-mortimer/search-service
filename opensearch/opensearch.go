package opensearch

import (
	"log"
	"os"

	"github.com/opensearch-project/opensearch-go"
)

var Client *opensearch.Client

func Init() {
	var err error
	Client, err = opensearch.NewClient(opensearch.Config{
		Addresses: []string{os.Getenv("OPENSEARCH_URL")},
	})
	if err != nil {
		log.Fatalf("Error creating OpenSearch client: %s", err)
	}
	log.Println("Connected to OpenSearch")
}
