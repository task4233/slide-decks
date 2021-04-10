package main

// error のメソッドセットはError
type error interface {
	Error() string
}
