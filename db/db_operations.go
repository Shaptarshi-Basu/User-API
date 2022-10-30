package db

import (
	"api-proj/model"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DBOperations interface {
	FetchUsers() ([]model.User, error)
	CreateUser(model.User) (model.User, error)
	UpdateUser(model.User) (model.User, error)
}

type DBOperationsImpl struct {
}

func createDynDBClient() *dynamodb.DynamoDB {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return dynamodb.New(session)
}
func (dbOps *DBOperationsImpl) FetchUsers() (users []model.User, err error) {
	fmt.Printf("List users")
	client := createDynDBClient()
	input := &dynamodb.ScanInput{
		TableName: aws.String("User"),
	}
	result, err := client.Scan(input)
	for _, i := range result.Items {
		user := model.User{}

		err = dynamodbattribute.UnmarshalMap(i, &user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (dbOps *DBOperationsImpl) CreateUser(user model.User) (model.User, error) {

	client := createDynDBClient()
	userData, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return model.User{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      userData,
		TableName: aws.String("User"),
	}
	_, err = client.PutItem(input)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
func (dbOps *DBOperationsImpl) UpdateUser(user model.User) (model.User, error) {
	fmt.Printf("Update user: %+v", user)
	client := createDynDBClient()
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":l": {
				S: aws.String(user.LastName),
			},
			":a": {
				S: aws.String(user.Address),
			},
		},
		TableName: aws.String("User"),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(user.Email),
			},
			"first_name": {
				S: aws.String(user.FirstName),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set last_name = :l, address = :a"),
	}
	_, err := client.UpdateItem(input)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
