package main

import (
	"fmt"
	"time"
)

func main() {
	day := "Sunday"
	switch day {
	case "Sunday", "Saturday":
		fmt.Println("Weekend! No work")
	case "Monday", "Tuesday":
		fmt.Println("Work days. Lots of meetings")
	default:
		fmt.Println("Mid-week")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good Morning")
	case hour < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}

	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Integer %d\n", v)
		case bool:
			fmt.Printf("Boolean %t\n", v)
		case string:
			fmt.Printf("String %s\n", v)
		case float32:
			fmt.Printf("Float %v", v)
		default:
			fmt.Printf("Unknown %T\n", v)
		}
	}
	checkType(30)
	checkType("String")
	checkType(true)
	checkType(324.9)
}
