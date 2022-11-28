package gen

import (
	"fmt"
)

// NumStr
// 2 泛型的类型限制，声明一个interface，包括所有需要支持的类型
// ~int 表示底层数据是int
type NumStr interface {
	~int | ~uint | ~float64 | ~string
}

type MyInt int

func AddNumStr[T NumStr](params []T) (sum T) {
	for _, param := range params {
		sum += param
	}
	return
}

func Test2() {
	// 2.1 支持的int
	intSum := AddNumStr([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Printf("2.1: 类型=%T，val=%+v\n", intSum, intSum)

	// 2.2 传入支持的string
	strSum := AddNumStr([]string{"哈", "哈", "哈", "哈"})
	fmt.Printf("2.2: 类型=%T，val=%+v\n", strSum, strSum)

	// 2.3 传入自定义int
	myIntSum := AddNumStr([]MyInt{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Printf("2.3: 类型=%T，val=%+v\n", myIntSum, myIntSum)
}
