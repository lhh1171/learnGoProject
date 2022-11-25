package myintf

import "fmt"

type ops func(int, int) int

func (p ops) Caller(a int, b int) int {
	fmt.Println("=============")
	return p(a, b)
}

type ops2 func(int, int) int

func (p ops2) Caller(a int, b int) int {
	fmt.Println("-----------------")
	return p(a, b)
}

func plus(a int, b int) int {
	return a + b
}
func TestOps() {

	fmt.Println(ops(plus).Caller(1, 2))
	fmt.Println(ops2(plus).Caller(1, 2))
}
