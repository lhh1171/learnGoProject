package gen

import (
	"fmt"
)

// M
// 4. 泛型map
type M[K string, V any] map[K]V

func Test4() {
	m := M[string, int]{
		"aaa": 123,
		"bbb": 456,
		"ccc": 789,
	}
	fmt.Printf("4.1: 类型=%T，val=%+v\n", m, m)
}
