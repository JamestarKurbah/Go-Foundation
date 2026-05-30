package main

import (
	"fmt"
	"os"
)

func simpleDefer() {
	fmt.Println("Defer Function simple defer")
	defer fmt.Println("Defer simple defer")
	fmt.Println("Defer Function Middle")

}
func lifoSimpleDefer() {
	fmt.Println("Funtion sinpleDefer: Start")
	defer fmt.Println("First: deffered")
	defer fmt.Println("Second: deffered")
	fmt.Println("Function lifoSimpleDefer: Middle")
}
func main() {

	file, err := os.Create("./defer.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	defer func() {
		fmt.Println("Before the return main")
	}()
	lifoSimpleDefer()
	fmt.Println("last line in main")
}
