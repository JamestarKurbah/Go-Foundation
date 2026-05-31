package main

import "fmt"

// use generic to make remove type restriction for our function
type Number interface {
	int | int8 | int16 | int64 | float64 | float32
}

func Sum[T Number](numbers ...T) T {
	var total T
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {

	grades := []int{90, 85}
	people := []string{"jane", "John", "mark"}
	fmt.Println(len(grades), len(people))

	fmt.Println(Sum[float64](10, 20, 30))
	v := Sum(10, 20, 30.9)
	fmt.Printf("%T\n", v)
}
