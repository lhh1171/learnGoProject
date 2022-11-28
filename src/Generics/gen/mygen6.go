package gen

import (
	"fmt"
	"strconv"
)

// FlyAnimal
// 6. 方法约束
type FlyAnimal interface {
	ToString() string
}

// Dragon
// Dragon实现了FlyAnimal
type Dragon int

func (d Dragon) ToString() string {
	return "string_" + strconv.Itoa(int(d))
}

// Tiger
// Tiger没有实现flyAnimal
type Tiger int

func PrintStr[T FlyAnimal](params ...T) {
	for _, param := range params {
		fmt.Println(param.ToString())
	}
}

func Test6() {
	// 6.1 传入实现了方法的类型
	dragon := Dragon(1)
	PrintStr(dragon)

	// 6.2 传入未实现对应方法的类型  ./generics_test.go:136:13: Tiger does not implement FlyAnimal (missing ToString method)
	//tiger := Tiger(100)
	//PrintStr(tiger)
}
