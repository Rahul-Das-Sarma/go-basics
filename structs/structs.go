package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := user{firstName, lastName, birthdate, time.Now()}

	outputPrint(appUser)
}

func outputPrint(u user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}
func getUserData(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scan(&input)
	return input
}
