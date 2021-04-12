type Num int

// 型 Num のメソッドセットは
// - addOneWithValueReceiver
// - addWithValueReceiver
func (num Num) addOneWithValueReceiver() { num++ }

func (num Num) addWithValueReceiver(val int) { num += val }

// 型 *Numのメソッドセットは
// - addOneWithPointerReceiver
// - addWithPointerReceiver
// - addOneWithValueReceiver
// - addWithValueReceiver
func (num *Num) addOneWithPointerReceiver() { *num++ }

func (num *Num) addWithPointerReceiver(val int) { *num += val }
