package gen

import (
	"fmt"
)

// Ch
// 5. 泛型通道
type Ch[T any] chan T

func Test5() {
	ch := make(Ch[int], 1)
	ch <- 10

	res := <-ch
	fmt.Printf("5.1: 类型=%T，val=%+v\n", res, res)
	fmt.Printf("5.2: 类型=%T，val=%+v\n", ch, ch)
}
