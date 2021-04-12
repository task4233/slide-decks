package main

func main() {
	var _ interface{} = nil

	var Num interface{} = -1

	// Numと1の型が異なるので、invalidな式の演算が原因でcompile errorになる
	// var _ int = Num + 1

	// intとNumのUnderlying typeが異なるので、Coversionできないことが原因でcompile errorになる
	// var _ int = int(Num) + 1

	// type assertionすればOK
	var _ int = Num.(int) + 1
}
