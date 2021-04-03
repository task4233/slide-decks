package main

import "fmt"

type Person struct{ Name string }

func (p Person) SayMorningByValue() { fmt.Printf("%sさん、おはよう！\n", p.Name) }

func (p *Person) SayMorningByPointer() { fmt.Printf("%sさん、おはよう！", p.Name) }

func main() {
	var p Person = Person{Name: "Gopher"}
	var pPointer *Person = &Person{Name: "Pointer Gopher"}

	p.SayMorningByValue()
	p.SayMorningByPointer()

	pPointer.SayMorningByValue()
	pPointer.SayMorningByPointer()
}
