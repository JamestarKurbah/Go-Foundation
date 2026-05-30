package main

import "fmt"

type Person interface {
	GetName() string
}

type Employee struct {
	ID   int
	Name string
}
type BusinessPerson struct {
	ID   int
	Name string
}

func (e Employee) GetName() string {
	return e.Name + " title"
}
func (e BusinessPerson) GetName() string {
	return e.Name
}
func displayName(p Person) {
	fmt.Println(p.GetName())
}
func main() {
	joe := Employee{ID: 1, Name: "Joe"}
	jane := BusinessPerson{ID: 1, Name: "Jane"}
	displayName(joe)
	displayName(jane)
}
