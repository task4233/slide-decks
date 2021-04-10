package main

type List []int

func (l List) Len() int { return len(l) }

func (l *List) Append(num int) { *l = append(*l, num) }
