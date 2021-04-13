package main

import "fmt"

type Num int

// Num型(defined type)のレシーバなので値レシーバ(value receiver)
func (num Num) addOneWithValueReceiver() { num++ }

// *Num型のレシーバなのでポインタレシーバ(pointer receiver)
func (num *Num) addOneWithPointerReceiver() { *num++ }

func main() {
	num := Num(2)
	num.addOneWithValue()
	fmt.Println(num)

	num.addOneWithPointer()
	fmt.Println(num)
}
