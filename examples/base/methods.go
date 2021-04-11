package main

import "fmt"

type Num int

func (num Num) addOneWithValue() { num++ }

func (num *Num) addOneWithPointer() { *num++ }

func main() {
	num := Num(2)
	num.addOneWithValue()
	fmt.Println(num)

	num.addOneWithPointer()
	fmt.Println(num)
}
