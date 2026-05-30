package main

import "fmt"

func main() {
	tmp := 25
	if tmp > 30 {
		fmt.Println("The weather is hot")
	} else {
		fmt.Println("The Weather is pleasant")
	}

	// map
	userAccess := map[string]bool{
		"jane": true,
		"john": false,
	}

	if hasAccess, ok := userAccess["john"]; ok && hasAccess {
		fmt.Println("Jane has access the system")
	} else {
		fmt.Println("access denied!")
	}
	// or
	hasAccess, ok := userAccess["jane"]
	if ok && hasAccess {
		fmt.Println("Jane has access the system")
	} else {
		fmt.Println("access denied!")
	}
}
