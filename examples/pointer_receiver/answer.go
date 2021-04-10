package main

import "fmt"

type error interface {
	Error() string
}

type EmptyError struct {
	FieldName string
}

func (e *EmptyError) Error() string {
	return fmt.Sprintf("%s is empty", e.FieldName)
}

func main() {
	// EmptyError型は、Error メソッドを実装していない
	// var _ error = EmptyError{}

	// 実装したいなら型を合わせる必要がある
	var _ error = &EmptyError{}
}
