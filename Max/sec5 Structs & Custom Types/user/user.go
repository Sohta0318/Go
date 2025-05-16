package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct{
	// FirstName string
	firstName string
	lastName string
	birthday string
	createdAt time.Time
}

type Admin struct{
  email string
  password string
  User
}

func (user *User) OutputUserDetails(){
	fmt.Println(user.firstName, user.lastName, user.birthday)
}

func (u *User) ClearUserName(){
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin{
  return Admin{
    email: email,
    password: password,
    User: User{
      firstName: "ADMIN",
      lastName: "ADMIN",
      birthday: "-----",
      createdAt: time.Now(),
    },
  }
}

func New (firstName, lastName, birthdate string) (*User, error){
	if firstName == "" || lastName == "" || birthdate == ""{
return nil, errors.New("First name, last name and birthdate are required")
	}
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthday: birthdate,
		createdAt: time.Now(),
	}, nil
}