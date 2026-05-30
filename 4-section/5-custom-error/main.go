package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrDivisionByZero = errors.New("Divide By Zero")
var ErrNumTooLarge = errors.New("Number too large")

type OpError struct {
	Op      string
	Code    int
	Message string
	Time    time.Time
}

func (op OpError) Error() string {
	return op.Message
}
func NewOpError(op string, code int, message string, t time.Time) *OpError {
	return &OpError{
		Op:      op,
		Code:    code,
		Message: message,
		Time:    t,
	}
}

func DoSomething() error {
	return NewOpError("doSomething", 100, "do something failed", time.Now())
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}

	if a > 1000 {
		return 0, ErrNumTooLarge
	}

	return a / b, nil
}

func main() {

	value, err := divide(4000, 1110)
	if err != nil {
		if err.Error() == ErrNumTooLarge.Error() {
			fmt.Println("do something else")
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println(value)
}
