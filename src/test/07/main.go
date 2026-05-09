package main

import "fmt"

type Pair struct {
	n, m int
}

func (p Pair) Combination() int {
	if p.m > p.n {
		return 0
	} else if p.m == 0 || p.m == p.n {
		return 1
	} else {
		return Pair{p.n - 1, p.m - 1}.Combination() + Pair{p.n - 1, p.m}.Combination()
	}
}

func main() {
	fmt.Println(Pair{5, 2}.Combination()) // 10
}
