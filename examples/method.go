package main

import "fmt"

type List []int

func (l List) Len() int { return len(l) }

// []intはdefined typeではないので、レシーバにはできない
// func (l []int) Len() int { return len(l)}

func main() {
	l := List{1, 3, 5}

	fmt.Println(l.Len())
	// 第一引数にレシーバを渡しても呼び出せる
	fmt.Println(List.Len(l))
}
