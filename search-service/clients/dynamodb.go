package clients

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var DynamoDBClient *dynamodb.Client

func InitializeDynamoDBClient(region string, customEndpoint string) {
	// Create a context
	ctx := context.TODO()

	// Load the default AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config: %v", err)
	}

	if customEndpoint != "" {
		log.Printf("Initializing DynamoDB client with Endpoint: %v", customEndpoint)
		DynamoDBClient = dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
			o.BaseEndpoint = aws.String(customEndpoint)
		})
		return
	}

	log.Print("Initializing DynamoDB client")
	// Initialize the DynamoDB client
	DynamoDBClient = dynamodb.NewFromConfig(cfg)
}

// GetDynamoDBClient returns the DynamoDB client
func GetDynamoDBClient() *dynamodb.Client {
	return DynamoDBClient
}
