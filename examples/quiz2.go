package main

import "fmt"

type Primes []int

func (p *Primes) Append(num int) []int {
	*p = append(*p, num)
	return *p
}

func main() {
	primes := (Primes{2, 3, 5}).Append(57)
	fmt.Println(primes)
}
