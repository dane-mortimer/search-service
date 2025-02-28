package clients

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	opensearch "github.com/opensearch-project/opensearch-go/v4"
	requestsigner "github.com/opensearch-project/opensearch-go/v4/signer/awsv2"
)

// OpenSearchClient is a wrapper for the OpenSearch client

var Client *opensearch.Client

// NewOpenSearchClient creates a new OpenSearch client based on the environment variable
func NewOpenSearchClient(env string, opensearchEndpoint string, awsRegion string) {

	var client *opensearch.Client
	var err error

	switch env {
	case "local":
		client, err = createLocalOpenSearchClient(opensearchEndpoint)
	default:
		client, err = createAWSOpenSearchClient(opensearchEndpoint, awsRegion)
	}

	if err != nil {
		fmt.Errorf("Failed to connect to Opensearch Endpoint")
		os.Exit(1)
	}

	Client = client
}

// createAWSOpenSearchClient creates an OpenSearch client for AWS OpenSearch with SigV4 signing
func createAWSOpenSearchClient(opensearchEndpoint string, awsRegion string) (*opensearch.Client, error) {
	// Load AWS configuration using the default credential chain
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion), // Use the AWS_REGION environment variable
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	// Create an AWS SigV4 signer using the loaded configuration
	signer, err := requestsigner.NewSignerWithService(cfg, "es")
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS SigV4 signer: %w", err)
	}

	// Create the OpenSearch client with AWS SigV4 signing
	client, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{opensearchEndpoint}, // AWS OpenSearch endpoint
		Signer:    signer,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create OpenSearch client: %w", err)
	}

	return client, nil
}

// createLocalOpenSearchClient creates an OpenSearch client for a local OpenSearch instance
func createLocalOpenSearchClient(opensearchEndpoint string) (*opensearch.Client, error) {
	// Create the OpenSearch client for local instance
	client, err := opensearch.NewClient(opensearch.Config{
		Addresses: []string{opensearchEndpoint}, // Local OpenSearch endpoint
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
