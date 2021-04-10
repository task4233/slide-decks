package main

func main() {
	var _ interface{} = nil

	var Num interface{} = -1

	// 型が異なるので実装できない
	// var _ int = Num + 1

	// Underlying typeが異なるのでCoversionできない
	// var _ int = int(Num) + 1

	// type assertionすればOK
	var _ int = Num.(int) + 1
}
