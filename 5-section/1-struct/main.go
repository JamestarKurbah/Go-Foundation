package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

func NewEmployee(id int, firstname, lastName, position string, salary int, isActive bool) Employee {
	return Employee{
		ID:        id,
		FirstName: firstname,
		LastName:  lastName,
		Position:  position,
		Salary:    salary,
		IsActive:  isActive,
		JoinedAt:  time.Now(),
	}
}

func main() {
	jane := Employee{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Position:  "Night",
		Salary:    10000,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}
	fmt.Println(jane)
	john := NewEmployee(2, "John", "Doe", "Day", 2000, true) // like constructor
	john.Salary = 3000
	fmt.Println(john)

	joePtr := &john
	fmt.Println(*joePtr)
	joePtr.IsActive = false
	joePtr.LastName = "John Adam"
	fmt.Println(john)
}
