package main

var _ error = &EmptyError{}

type error interface {
	Error() string
}

type EmptyError struct{}

// Error は型 *EmptyError のメソッドセットにも含まれる
func (e EmptyError) Error() string {
	return "empty"
}

func main() {}
