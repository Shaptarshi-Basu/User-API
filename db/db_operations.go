package db

import (
	"user-api/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DBOperations interface {
	FetchUsers() ([]model.User, error)
	CreateOrUpdateUser(model.User) (model.User, error)
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
	users = []model.User{}
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
func (dbOps *DBOperationsImpl) CreateOrUpdateUser(user model.User) (model.User, error) {
	client := createDynDBClient()
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":f": {
				S: aws.String(user.FirstName),
			},
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
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set first_name = :f, last_name = :l, address = :a"),
	}
	_, err := client.UpdateItem(input)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
