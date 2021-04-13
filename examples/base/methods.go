package main

import "fmt"

type Num int

// List型のレシーバなので値レシーバ
func (l List) AppendWithValueReceiver(num int) { l = append(l, num) }

// *List型のレシーバなのでポインタレシーバ
func (l *List) AppendWithPointerReceiver(num int) { *l = append(*l, num) }

func main() {
	num := Num(2)
	num.addOneWithValue()
	fmt.Println(num)

	num.addOneWithPointer()
	fmt.Println(num)
}
