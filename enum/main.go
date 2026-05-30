package main

import "fmt"

const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

type LogLevel int

const (
	LogError LogLevel = iota
	LogWarn           = 1
	LogInfo           = 2
	LogDebug          = 3
	LogFatal          = 4
)

func main() {
	fmt.Println(Monday)
	fmt.Println(Friday)
}
