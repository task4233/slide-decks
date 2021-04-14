package main

var _ error = EmptyError{}

type error interface {
	Error() string
}

type EmptyError struct{}

func (e EmptyError) Error() string {
	return "error"
}

func main() {}
