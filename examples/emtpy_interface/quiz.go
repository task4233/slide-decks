package main

type EmptyInterface interface{}

func main() {
	var _ EmptyInterface = nil
}
