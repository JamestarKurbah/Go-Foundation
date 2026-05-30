package main

import "fmt"

type Person interface {
	GetName() string
}

type BusinessPerson struct {
	ID   int
	Name string
}

type ID int

func (idx ID) String() string {
	return fmt.Sprintf("ID: %d", idx)
}

func (e BusinessPerson) GetName() string {
	return e.Name
}
func (e BusinessPerson) String() string {
	return fmt.Sprintf("Person ID[:%d, Name:%s]", e.ID, e.Name)
}
func main() {
	jane := BusinessPerson{ID: 1, Name: "Jane"}
	fmt.Println(jane)
	myID := ID(40)
	fmt.Println(myID)
}
