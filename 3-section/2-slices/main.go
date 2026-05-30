package main

import "fmt"

func main() {

	names := []string{"Alice", "John", "Mark"}
	items := make([]int, 3, 5)
	fmt.Printf("%+v\n", names)

	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))
	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 5)
	items = append(items, 8)
	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 5)
	items = append(items, 8)
	// items = append(items, 1)
	// items = append(items, 2)
	// items = append(items, 5)
	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))

	// in a slice have len and capacity if data exceed capacity go will double the capacity and copy data to new location

	fmt.Printf("%+v", items[3:7])
}
