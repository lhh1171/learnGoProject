package gen

import (
	"fmt"
)

// Vector
// 3. 泛型切片
// any为泛型自带的一种类型，即interface
type Vector[T any] []T
type NumSlice[T int | float64] []T

func Test3() {
	// 测试3.1
	v := Vector[string]{"z", "x", "c"}
	fmt.Printf("3.1: 类型=%T，val=%+v\n", v, v)

	// 测试3.2
	ns := NumSlice[int]{1, 2, 3, 4, 5, 6}
	fmt.Printf("3.2: 类型=%T，val=%+v\n", ns, ns)

	// 测试3.3
	sum := AddElem(ns)
	fmt.Printf("3.3: 类型=%T，val=%+v\n", sum, sum)
}
