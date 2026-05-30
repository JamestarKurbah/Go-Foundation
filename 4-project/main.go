package main

import "fmt"

type LogLevel int

const (
	LevelTrace   LogLevel = iota
	LevelDebug            = 1
	LevelInfo             = 2
	LevelWarning          = 3
	levelError            = 4
)

var levelNames = []string{"Trace", "Debug", "Info", "Warning", "Error"}

func (l LogLevel) String() string {
	if l < LevelTrace || l > levelError {
		return "unknown"
	}
	return levelNames[l]
}

func printLogLevel(level LogLevel) {
	fmt.Printf("Log Level: %d %s\n", level, level.String())
}

func main() {
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
	printLogLevel(levelError)
	printLogLevel(10)
}
