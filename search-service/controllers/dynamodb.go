package controllers

import (
	"context"
	"os"
	"search-service/clients"
	"search-service/dao"
	"search-service/models"
)

func GetCourseController(id string) (models.Course, error) {

	ctx := context.TODO()

	client := clients.GetDynamoDBClient()

	tableName := os.Getenv("COURSE_TABLE")

	courseTable := dao.CourseTable{
		DynamoDbClient: client,
		TableName:      tableName,
	}

	course, err := courseTable.GetCourse(ctx, id)

	return course, err
}

func CreateCourseController(course models.Course) (*models.Course, error) {

	ctx := context.TODO()

	client := clients.GetDynamoDBClient()

	tableName := os.Getenv("COURSE_TABLE")

	courseTable := dao.CourseTable{
		DynamoDbClient: client,
		TableName:      tableName,
	}

	result, err := courseTable.CreateCourse(ctx, course)
	if err != nil {
		return nil, err
	}

	return result, nil
}
