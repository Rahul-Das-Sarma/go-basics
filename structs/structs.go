package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name, and birthdate are required")
	}
	return &user{firstName, lastName, birthdate, time.Now()}, nil
}
func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := newUser(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputPrint(appUser)
	appUser.clearUser()

	fmt.Println("checking:", appUser.firstName, appUser.lastName, appUser.birthdate)
}
func (u *user) clearUser() {
	u.firstName = ""
	u.lastName = ""

}
func outputPrint(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}
func getUserData(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}
