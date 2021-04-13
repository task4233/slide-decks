package main

import "fmt"

var _ error = &EmptyError{}

type error interface {
	Error() string
}

type EmptyError struct {
	FieldName string
}

func (e *EmptyError) Error() string {
	return fmt.Sprintf("%s is empty", e.FieldName)
}

func main() {}
