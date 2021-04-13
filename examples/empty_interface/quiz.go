package main

type Empty interface{}

func main() {
	var _ Empty = nil
}
