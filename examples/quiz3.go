package main

import "fmt"

type Primes []int

var primes Primes

func (p Primes) AddressIsEqual() bool {
	return primes == &p
}

func main() {
	primes = &Primes{2, 3, 5}
	fmt.Println(primes.AddressIsEqual())
}
