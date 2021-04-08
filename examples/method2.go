package main

import "fmt"

type List []int

func (l List) AppendWithValueReceiver(num int) { l = append(l, num) }

func (l *List) AppendWithPointerReceiver(num int) { *l = append(*l, num) }

func main() {
	l := List{1, 3, 5}
	l.AppendWithValueReceiver(7)
	fmt.Println(l)

	l.AppendWithPointerReceiver(7)
	fmt.Println(l)
}
