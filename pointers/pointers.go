package main

import "fmt"

// without pointers

// func main() {
// 	age := 30
// 	fmt.Println(age)
// 	adultAge := getAdultAge(age)
// 	fmt.Println(adultAge)
// }
// func getAdultAge(age int) int {
// 	return age - 18
// }

// with pointers

func main() {
	age := 30

	var agePointer *int = &age

	fmt.Println(agePointer)
	adultAge := getAdultAge(agePointer)
	fmt.Println(adultAge)
}

func getAdultAge(age *int) int {
	return *age - 18
}
