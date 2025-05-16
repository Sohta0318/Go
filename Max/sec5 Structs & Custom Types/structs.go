package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	var appUser *user.User

	appUser, err := user.New(firstName,lastName, birthdate)

	if err != nil{
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@email.com", "test123")

	admin.ClearUserName()
	admin.OutputUserDetails()

	// outputUserDetails(&appUser)

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

// func outputUserDetails(user *user){
// 	// ...
// 	// (*u).firstName
// 	fmt.Println(user.firstName, user.lastName, user.birthday)

// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
