package main

import "fmt"

func main() {
	var numbers [2]int
	fmt.Printf("%+v\n", numbers)
	numbers[0] = 1
	numbers[1] = 2
	fmt.Printf("%+v\n", numbers)
	primes := [4]int{2, 3, 5, 7}
	fmt.Printf("%+v\n", primes)
	primes[3] = 11
	fmt.Printf("%+v\n", primes)
	for i := 0; i < len(primes); i++ {
		fmt.Printf("%+v\n", primes[i])
	}

	var matix [2][3]int
	matix[0][0] = 1
	matix[0][1] = 2
	matix[0][2] = 3
	matix[1][0] = 4
	matix[1][1] = 5
	matix[1][2] = 6
	fmt.Printf("%+v", matix)

	// slice

}
