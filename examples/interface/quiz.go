package main

type Animal interface {
	MakeSound() string
}

type Cat struct {
}

func (Cat) MakeSound() []byte {
	return []byte("meow")
}

func main() {
	// Cat構造体はAnimalインタフェースを実装している?
	var _ Animal = Cat{}
}
