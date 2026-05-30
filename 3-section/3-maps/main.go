package main

import "fmt"

func main() {
	studentGrades := map[string]int{
		"Alice": 90,
		"James": 85,
		"Dan":   60,
	}
	fmt.Printf("%+v\n", studentGrades)
	studentGrades["Alice"] = 96
	fmt.Printf("%+v\n", studentGrades)
	alice, ok := studentGrades["Alice"]
	if ok {
		fmt.Printf("Alice: %+v\n", alice)
	}
	delete(studentGrades, "Alice")
	fmt.Printf("%+v\n", studentGrades)

	key := "John"
	if value, ok := studentGrades[key]; ok {
		fmt.Printf("%s: %+v\n", key, value)
	}
	configs := make(map[string]int)
	fmt.Printf("\n%+v %T\n", configs, configs)

	// var test map[string]int just initial without assign map will be nil
	test := map[string]int{} //here map is declare and assign some empty
	if test == nil {
		fmt.Println("Test is nil")
	}

}
