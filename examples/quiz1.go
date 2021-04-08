package main

import "fmt"

type Primes []int

func (p Primes) Len() int {
	return len(p)
}

func main() {
	length := (&Primes{2, 3, 5}).Len()
	fmt.Println(length)
}
