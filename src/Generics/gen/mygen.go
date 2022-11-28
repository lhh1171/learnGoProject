package gen

import (
	"fmt"
)

// AddElem
// 1. 泛型的类型限制，在函数上直接申明该函数支持的多个类型
func AddElem[T int | string](params []T) (sum T) {
	for _, elem := range params {
		sum += elem
	}
	return
}

func Test1() {
	// 1. 在函数上声明泛型支持的多个类型
	// 1.1 传入支持的int
	intSum := AddElem([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Printf("1.1: 类型=%T，val=%+v\n", intSum, intSum)
	// 1.2 传入支持的string
	strSum := AddElem([]string{"哈", "哈", "哈", "哈", "，", "哈", "哈", "哈", "哈"})
	fmt.Printf("1.2: 类型=%T，val=%+v\n", strSum, strSum)
	// 1.3 传入不支持的类型  ./generics_test.go:29:24: float64 does not implement int|string
	//floatSum := AddElem([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	//fmt.Printf("1.3: 类型=%T，val=%+v\n", floatSum, floatSum)
}
