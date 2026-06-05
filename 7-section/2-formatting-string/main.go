package main

import (
	"errors"
	"fmt"
)

type ConfigItem struct {
	Key   string
	Value interface{}
	IsSet bool
}

/*
%v	- the default formatting
%+v - you can see the name of the feed
%#v - go syntax representation value
%T  - to print type
%s  - string
%d  - integer
%f  - float(%.2f)
%t  - boolean
%q  - double code string
%%  - printing %
*/
func (c ConfigItem) String() string {
	return fmt.Sprintf("Key:%s, Value: %v, IsSet: %t", c.Key, c.Value, c.IsSet)
}
func main() {
	appName := "envParser"
	version := 1.2
	port := 8080
	isEnabled := true
	status := fmt.Sprintf("Application: %s (version: %.1f) running on port %d. Enabled: %t", appName, version, port, isEnabled)
	fmt.Println(status)
	item1 := ConfigItem{Key: "API_URL", Value: "html://localhost:3000/api", IsSet: true}
	item2 := ConfigItem{Key: "TIMEOUT_MS", Value: 5000, IsSet: true}
	item3 := ConfigItem{Key: "DEBUG_MODE", Value: false, IsSet: false}
	fmt.Printf("Item 1 (%%v): %v\n", item1)
	fmt.Printf("Item 2 (%%+v): %v\n", item2)
	fmt.Printf("Item 3 (%%#v): %v\n", item3)
	err := errors.New("Test")
	fmt.Errorf("here is the error on port %d: %w", port, err)

}
