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
	// Cat構造体はAnimalインタフェースを実装している
	var _ Animal = Cat{}
}
