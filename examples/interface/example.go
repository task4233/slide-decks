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
	// Cat型はAnimalインタフェースを実装している
	var _ Animal = Cat{}
}
