package main

import "fmt"

type error interface {
	Error() string
}

type EmptyError struct {
	FieldName string
}

func (e EmptyError) Error() string {
	return fmt.Sprintf("%s is not found", e.FieldName)
}

func main() {
	var _ error = EmptyError{}
}
