package main

import "fmt"

// addOne はfunction(関数)
func addOne(num int) int {
	return num + 1
}

func main() {
	var primeNum int = 57
	fmt.Println(addOne(primeNum))
}
