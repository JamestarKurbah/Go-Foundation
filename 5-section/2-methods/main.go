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

// FullName pointer receiver
func (e *Employee) FullName() string {
	if e == nil {
		return ""
	}
	return e.FirstName + " " + e.LastName
}

// // Deactivate value receiver
// func (e Employee) Deactivate() { // copy of data
// 	e.IsActive = false
// }

// Deactivate pointer receiver
func (e *Employee) Deactivate() { // copy of data
	e.IsActive = false
}

func (e *Employee) SetJoinedAt(t time.Time) {
	e.JoinedAt = t
}

func main() {
	jane := Employee{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Position:  "Night",
		Salary:    10000,
		IsActive:  true,
		//JoinedAt:  time.Now(),
	}
	fmt.Println(jane)
	jane.Deactivate()
	fmt.Println(jane)
	jane.SetJoinedAt(time.Now().Add(100000000 * time.Minute))
	fmt.Println(jane)
	// demo to show panic
	joe := &Employee{}
	joe = nil
	joe.FullName()
}
