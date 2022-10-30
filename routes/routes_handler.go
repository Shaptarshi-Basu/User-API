package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"user-api/db"
	"user-api/model"

	"github.com/aws/aws-lambda-go/events"
)

type OpertionHandlers struct {
	dbOps db.DBOperations
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	dbOps := db.DBOperationsImpl{}
	oh := OpertionHandlers{dbOps: &dbOps}
	if req.HTTPMethod == "GET" && req.Path == "/users" {
		return oh.fetchAllUsers()
	} else if req.HTTPMethod == "POST" && req.Path == "/user" {
		return oh.createUser(req)
	} else if req.HTTPMethod == "PUT" && req.Path == "/user" {
		return oh.updateUser(req)
	} else {
		return clientError(http.StatusBadRequest, fmt.Errorf("Invalid Route, Method: %s and Path: %s", req.HTTPMethod, req.Path))
	}
}

func (oh *OpertionHandlers) fetchAllUsers() (events.APIGatewayProxyResponse, error) {
	users, err := oh.dbOps.FetchUsers()
	if err != nil {
		return clientError(http.StatusInternalServerError, err)
	}

	usersFetched, err := json.Marshal(users)
	if err != nil {
		return clientError(http.StatusInternalServerError, err)
	}

	return returnResponse(string(usersFetched))
}
func (oh *OpertionHandlers) createUser(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var userDetails model.User
	err := json.Unmarshal([]byte(req.Body), &userDetails)
	if err != nil {
		return clientError(http.StatusInternalServerError, err)
	}
	fmt.Printf("%+v", userDetails)
	err = validateUserDetails(userDetails)
	if err != nil {
		return clientError(http.StatusBadRequest, err)
	}

	user, err := oh.dbOps.CreateUser(userDetails)
	if err != nil {
		return clientError(http.StatusInternalServerError, err)
	}

	userCreated, err := json.Marshal(user)
	if err != nil {
		return clientError(http.StatusInternalServerError, err)
	}

	return returnResponse(string(userCreated))
}

func (oh *OpertionHandlers) updateUser(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var userDetails model.User
	err := json.Unmarshal([]byte(req.Body), &userDetails)
	if err != nil {
		return clientError(http.StatusBadRequest, err)
	}
	err = validateUserDetails(userDetails)
	if err != nil {
		return clientError(http.StatusBadRequest, err)
	}

	user, err := oh.dbOps.UpdateUser(userDetails)
	if err != nil {
		return clientError(http.StatusInternalServerError, err)
	}

	userUpdated, err := json.Marshal(user)
	if err != nil {
		return clientError(http.StatusInternalServerError, err)
	}

	return returnResponse(string(userUpdated))
}

func validateUserDetails(user model.User) (err error) {
	if user.Email == "" {
		err = fmt.Errorf("Email is a required field and cannot be emoty in the request")
	} else if user.FirstName == "" {
		err = fmt.Errorf("FirstName is a required field and cannot be emoty in the request")
	}
	return err
}

func clientError(status int, err error) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Error: %+v", err)
	return events.APIGatewayProxyResponse{
		Body:       err.Error(),
		StatusCode: status,
	}, err
}
func returnResponse(body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: http.StatusOK,
	}, nil
}
