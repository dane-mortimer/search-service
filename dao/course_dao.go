package dao

import (
	"context"
	"log"
	"search-service/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type CourseTable struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func (courseTable CourseTable) GetCourse(ctx context.Context, id int) (models.Course, error) {
	course := models.Course{}
	response, err := courseTable.DynamoDbClient.GetItem(ctx, &dynamodb.GetItemInput{
		Key:       map[string]types.AttributeValue{"id": id},
		TableName: aws.String(courseTable.TableName),
	})
	if err != nil {
		log.Printf("Couldn't get course with ID %v. Here's why: %v\n", id, err)
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &course)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}
	return course, err
}
