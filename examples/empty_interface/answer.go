package main

type Empty interface{}

func main() {
	// 全ての値を代入可能
	var _ Empty = nil
	var _ Empty = 57
	var _ Empty = "hoge"

	type Person struct {
		Name string
	}
	var _ Empty = Person{}
}
