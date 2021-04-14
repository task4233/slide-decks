package main

import "fmt"

type EmptyError struct {
	FieldName string
}

func (e *EmptyError) Error() string {
	return fmt.Sprintf("%s is empty", e.FieldName)
}

func main() {
	var emptyError EmptyError = EmptyError{FieldName: "hoge"}

	// (&emptyError).Error() と同義
	fmt.Println(emptyError.Error())
}
