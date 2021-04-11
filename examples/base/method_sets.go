package main

type List []int

// 型 List のメソッドセットはAppendWithValueReceiver
func (l List) AppendWithValueReceiver(num int) { l = append(l, num) }

// 型 *List のメソッドセットはAppendWithPointerReceiverとAppendWithValueReceiver
func (l *List) AppendWithPointerReceiver(num int) { *l = append(*l, num) }

func main() {
}
