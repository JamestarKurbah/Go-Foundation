package main

import "fmt"

// panic recover
func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("something wrong in mightpanic")
	}
	fmt.Println("This function executed without panic")
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic: ", r)
		}
	}()
	mightPanic(true)
}
func main() {
	recoverable()
}
