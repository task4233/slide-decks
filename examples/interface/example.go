package main

type Animal interface {
	MakeSound() string
}

type Cat struct {
}

func (Cat) MakeSound() string {
	return "meow"
}

func main() {
	// Cat型はAnimal interfaceを実装している
	var _ Animal = Cat{}
}
