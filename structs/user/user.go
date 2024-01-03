package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate time.Time
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: time.Now(),
			createdAt: time.Now(),
		},
	}
}

// Receiver
func (user User) OutputUserDetails() {
	fmt.Println(user.firstName, user.lastName, user.birthdate, user.createdAt)
}

// Receiver with pointer to update a struct data
func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

// Constructor function for struct
func New(firstName string, lastName string, birthdate time.Time) (*User, error) {
	if firstName == "" || lastName == "" || birthdate.IsZero() {
		return nil, errors.New("invalid data")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
