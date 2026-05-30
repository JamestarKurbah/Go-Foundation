package main

import "fmt"

func modifyValue(val int) {
	val = val * 10
	fmt.Printf("val %d\n", val)
}
func modifyPointer(val *int) {
	if val == nil {
		fmt.Println("val is nil")
		return
	}
	*val = *val * 10 //deferrencing
	fmt.Printf("modifyPointer: %+v\n", val)
}
func main() {
	age := 10
	agePtr := &age
	fmt.Printf("age %d\n", agePtr)
	fmt.Printf("age %d\n", &age)
	modifyValue(age)
	fmt.Printf("%+v\n", &age)
	modifyPointer(agePtr)
	fmt.Printf("%+v\n", age)
	grade := 50
	gradePtr := &grade
	fmt.Printf("gradePtr grade: %+v\n", gradePtr)
	fmt.Printf("gradePtr: %+v\n", *(&gradePtr))

}
