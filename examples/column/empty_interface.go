package main

func main() {
	var _ interface{} = nil

	var Num interface{} = -1

	// Numと1の型が異なるのでinvalidな式の評価になる
	// var _ int = Num + 1

	// intとNumのUnderlying typeが異なるのでCoversionできない
	// var _ int = int(Num) + 1

	// type assertionすればOK
	var _ int = Num.(int) + 1
}
