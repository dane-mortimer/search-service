package clients

import (
	"context"
	"fmt"
	"search-service/middleware"
	"search-service/models"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func InitializeCognitoClient(userPoolID, region string, adminPaths []models.AdminPath) (*middleware.CognitoClient, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))

	if err != nil {
		return nil, fmt.Errorf("unable to laod SDK config: %v", err)
	}

	cognitoClient := cognitoidentityprovider.NewFromConfig(cfg)

	return &middleware.CognitoClient{
		UserPoolID:     userPoolID,
		Region:         region,
		CognitoClient:  cognitoClient,
		AdminOnlyPaths: adminPaths,
	}, nil
}
