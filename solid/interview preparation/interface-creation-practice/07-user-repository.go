// Repository Pattern
// Create a UserRepository interface with methods like Create, GetByID, Delete. Implement an in-memory and a mock version for testing

package interface_creation_practice

import "fmt"

type User struct {
	Name  string
	Bio   string
	Email string
}

type UserRepository interface {
	GetByID()
	Create()
	Delete()
}

type UserClient struct{}
type UserMock struct{}

func (uc UserClient) Create() {
	fmt.Print("Creating user at DB level")
}

func (um UserMock) Create() {
	fmt.Print("Creating user in mock mode")
}

func (uc UserClient) GetByID() {
	fmt.Print("Get By ID user at DB level")
}

func (um UserMock) GetByID() {
	fmt.Print("Get By ID user in mock mode")
}

func (uc UserClient) Delete() {
	fmt.Print("Deleting user at DB level")
}

func (um UserMock) Delete() {
	fmt.Print("Deleting user in mock mode")
}
