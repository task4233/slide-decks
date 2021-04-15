package main

type error interface {
	Error() string
}

type EmptyError struct{}

func (e *EmptyError) Error() string {
	return "empty"
}

func main() {
	var emptyError EmptyError
	// 内部的に(&emptyError).Error() と同義
	emptyError.Error()
}
