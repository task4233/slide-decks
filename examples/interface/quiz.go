package main

var _ Animal = Cat{}

type Animal interface {
	MakeSound() string
}

type Cat struct {
}

func (Cat) MakeSound() []byte {
	return []byte("meow")
}

func main() {}
