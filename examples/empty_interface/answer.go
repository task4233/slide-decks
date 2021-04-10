package main

type EmptyInterface interface{}

func main() {
	// 全ての値を代入可能
	var _ EmptyInterface = nil
	var _ EmptyInterface = 57
	var _ EmptyInterface = "hoge"

	type Person struct {
		Name string
	}
	var _ EmptyInterface = Person{}
}
