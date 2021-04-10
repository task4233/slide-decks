package main

import "fmt"

type error interface {
	Error() string
}

type EmptyError struct {
	FieldName string
}

func (e EmptyError) Error() string {
	return fmt.Sprintf("%s is empty", e.FieldName)
}

func main() {
	// EmptyError型は、Error メソッドを実装している？
	var _ error = &EmptyError{}
}
