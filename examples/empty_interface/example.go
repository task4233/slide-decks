package main

import "fmt"

var _ EmptyInterface = T{}

type EmptyInterface interface{}

type T struct{}

func (T) Hello() {
	fmt.Println("Hello!")
}

func main() {}
