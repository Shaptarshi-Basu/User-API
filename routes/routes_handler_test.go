package routes

import (
	mock "api-proj/mocks"
	"api-proj/model"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	dbOps := mock.DBOperationsMock{}
	oh := OpertionHandlers{dbOps: &dbOps}
	testUser := model.User{Email: "mockmail@gmail.com", FirstName: "SomeName", LastName: "SomeLAstName", Address: "SomeAddress"}
	reqbody, _ := json.Marshal(testUser)
	res, err := oh.createUser(events.APIGatewayProxyRequest{Body: string(reqbody)})
	var expectedUser model.User
	json.Unmarshal([]byte(res.Body), &expectedUser)
	assert.Equal(t, testUser, expectedUser)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Nil(t, err)
}

func TestCreateUserFail_Mandatory_Params_Missing(t *testing.T) {
	dbOps := mock.DBOperationsMock{}
	oh := OpertionHandlers{dbOps: &dbOps}
	testUser := model.User{Email: "", FirstName: "SomeName", LastName: "SomeLAstName", Address: "SomeAddress"}
	reqbody, _ := json.Marshal(testUser)
	res, err := oh.createUser(events.APIGatewayProxyRequest{Body: string(reqbody)})
	var expectedUser model.User
	json.Unmarshal([]byte(res.Body), &expectedUser)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("Email is a required field and cannot be emoty in the request"), err)
}

func TestCreateUserFail_DB_operations_Failure(t *testing.T) {
	dbOps := mock.DBOperationsFailure{}
	oh := OpertionHandlers{dbOps: &dbOps}
	testUser := model.User{Email: "someEmail@gmail.com", FirstName: "SomeName", LastName: "SomeLAstName", Address: "SomeAddress"}
	reqbody, _ := json.Marshal(testUser)
	res, err := oh.createUser(events.APIGatewayProxyRequest{Body: string(reqbody)})
	var expectedUser model.User
	json.Unmarshal([]byte(res.Body), &expectedUser)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("Couldn't create user"), err)
}
func TestUpdateUser(t *testing.T) {
	dbOps := mock.DBOperationsMock{}
	oh := OpertionHandlers{dbOps: &dbOps}
	testUser := model.User{Email: "mockmail@gmail.com", FirstName: "SomeName", LastName: "SomeLAstName", Address: "SomeAddress"}
	reqbody, _ := json.Marshal(testUser)
	res, err := oh.updateUser(events.APIGatewayProxyRequest{Body: string(reqbody)})
	var expectedUser model.User
	json.Unmarshal([]byte(res.Body), &expectedUser)
	assert.Equal(t, testUser, expectedUser)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Nil(t, err)
}

func TestUpdateUserFail_Mandatory_Params_Missing(t *testing.T) {
	dbOps := mock.DBOperationsMock{}
	oh := OpertionHandlers{dbOps: &dbOps}
	testUser := model.User{Email: "", FirstName: "SomeName", LastName: "SomeLAstName", Address: "SomeAddress"}
	reqbody, _ := json.Marshal(testUser)
	res, err := oh.updateUser(events.APIGatewayProxyRequest{Body: string(reqbody)})
	var expectedUser model.User
	json.Unmarshal([]byte(res.Body), &expectedUser)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("Email is a required field and cannot be emoty in the request"), err)
}

func TestUpdateUserFail_DB_operations_Failure(t *testing.T) {
	dbOps := mock.DBOperationsFailure{}
	oh := OpertionHandlers{dbOps: &dbOps}
	testUser := model.User{Email: "someEmail@gmail.com", FirstName: "SomeName", LastName: "SomeLAstName", Address: "SomeAddress"}
	reqbody, _ := json.Marshal(testUser)
	res, err := oh.updateUser(events.APIGatewayProxyRequest{Body: string(reqbody)})
	var expectedUser model.User
	json.Unmarshal([]byte(res.Body), &expectedUser)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("Couldn't update user"), err)
}
func TestFetchAllUsersSuccessFully(t *testing.T) {
	dbOps := mock.DBOperationsMock{}
	oh := OpertionHandlers{dbOps: &dbOps}
	res, err := oh.fetchAllUsers()
	var expectedUsrs []model.User
	testUser := model.User{Email: "mockmail@gmail.com", FirstName: "SomeName", LastName: "SomeLAstName", Address: "SomeAddress"}
	json.Unmarshal([]byte(res.Body), &expectedUsrs)
	assert.Equal(t, []model.User{testUser}, expectedUsrs)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Nil(t, err)
}
