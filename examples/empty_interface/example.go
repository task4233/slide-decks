package main

import "fmt"

type EmptyInterface interface {
}

type T struct {
}

func (T) Hello() {
	fmt.Println("Hello!")
}

func main() {
	// EmptyInterface のメソッドセットはなく、Tのメソッド
	var _ EmptyInterface = T{}
}
