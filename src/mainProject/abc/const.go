package abc

import (
	"fmt"
)

const (
	i1 int = 1
	i2
	i3
	i4 int = iota
	i5
)

const item1 = 1
const item2 = 2

func TestConst() {
	fmt.Println(item1, item2)
	fmt.Println(i1, i2, i3, i4, i5)
}
