package dao

import (
	"context"
	"log"
	"math/rand"
	"search-service/models"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/oklog/ulid"
)

type CourseTable struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func (courseTable CourseTable) CreateCourse(ctx context.Context, course models.Course) (*models.Course, error) {
	// Generate a ULID for the course ID
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	createdAt := time.Now().Format(time.RFC3339) // Use RFC3339 for ISO 8601 format
	course.CreatedAt = createdAt
	course.ID = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	// Marshal the course into a DynamoDB attribute map
	item, err := attributevalue.MarshalMap(course)
	if err != nil {
		log.Printf("Couldn't marshal course %v. Here's why: %v\n", course, err)
		return nil, err
	}

	_, err = courseTable.DynamoDbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(courseTable.TableName),
		Item:      item,
	})

	if err != nil {
		log.Printf("Couldn't insert course %v into table %v. Here's why: %v\n", course, courseTable.TableName, err)
	}

	log.Printf("Course %v inserted successfully into table %v.\n", course, courseTable.TableName)

	return &course, err
}

func (courseTable CourseTable) GetCourse(ctx context.Context, id string) (models.Course, error) {
	course := models.Course{}

	key := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: id}, // Use AttributeValueMemberS for string IDs
	}

	response, err := courseTable.DynamoDbClient.GetItem(ctx, &dynamodb.GetItemInput{
		Key:       key,
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
