package main

import (
	"fmt"
	"time"

	"github.com/dreking/structs/user"
)

// Custom type
type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	var name str = "Dre"
	name.log()

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdateString := getUserData("Please enter your birthdate (MM-DD-YYYY): ")
	fmt.Println(birthdateString)

	birthdate, err := time.Parse("01-02-2006", birthdateString)
	if err != nil {
		fmt.Println("Invalid birthdate format:", err)
		return
	}

	appUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@test.com", "test")
	admin.OutputUserDetails()
	admin.ClearUserName()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
