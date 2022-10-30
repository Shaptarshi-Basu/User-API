package mock

import (
	"api-proj/model"
	"fmt"
)

type DBOperationsMock struct {
}
type DBOperationsFailure struct {
}

func (dbOps *DBOperationsMock) FetchUsers() (users []model.User, err error) {
	fmt.Printf("List users")
	return []model.User{{Email: "mockmail@gmail.com", FirstName: "SomeName", LastName: "SomeLAstName", Address: "SomeAddress"}}, nil
}
func (dbOps *DBOperationsMock) UpdateUser(user model.User) (model.User, error) {
	return user, nil
}
func (dbOps *DBOperationsMock) CreateUser(user model.User) (model.User, error) {
	return user, nil
}

func (dbOpsFail *DBOperationsFailure) FetchUsers() (users []model.User, err error) {
	fmt.Printf("List users")
	return []model.User{}, fmt.Errorf("Couldnt Fetch User")
}
func (dbOpsFail *DBOperationsFailure) UpdateUser(user model.User) (model.User, error) {
	return model.User{}, fmt.Errorf("Couldn't update user")
}
func (dbOpsFail *DBOperationsFailure) CreateUser(user model.User) (model.User, error) {
	return model.User{}, fmt.Errorf("Couldn't create user")
}
