package main

var _ Animal = Cat{}

type Animal interface {
	MakeSound() string
}

type Cat struct {
}

func (Cat) MakeSound() string {
	return "meow"
}

func main() {}
