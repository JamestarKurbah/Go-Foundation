package main

import (
	"errors"
	"fmt"
	"strings"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]
	return parts[1], parts[0]
}

func main() {
	value, err := divide(4, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
	firstName, lastName := splitName("John Wick")
	fmt.Println("FirstName: ", firstName)
	fmt.Println("LastName: ", lastName)
}
