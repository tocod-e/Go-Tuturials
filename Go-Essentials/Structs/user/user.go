package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	age       int
	createdAt time.Time
}


func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last NAme and Birthdate are reqeuired")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func OutputUserDetails(u *User){
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

type Admin struct {
	email string
	password string
	User
}

func NewAdmin(email, password string) Admin{

	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName: "ADMIN",
			birthDate: "-----",
			createdAt: time.Now(),
		},
	}
}