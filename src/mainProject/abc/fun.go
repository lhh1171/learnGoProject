package abc

import (
	"fmt"
	"unsafe"
)

func Ff1() {
	plus := func(a, b int64) int64 {
		return a + b
	}
	ops := func(f func(int64, int64) int64, a, b int64) int64 {
		return f(a, b)
	}
	fmt.Println(ops(plus, 4, 2))
}
func Ff2() {
	plus := func(a, b int64) int64 {
		fmt.Println("anc")
		return a + b
	}
	type func2 func(int64, int64) int64
	ptr := new(func2)     //&func2
	ptr = (*func2)(&plus) // 指针类型都一样，*的时候不一样
	fmt.Println((*ptr)(3, 2))
	fmt.Println("大小为", unsafe.Sizeof(plus))
}
func PartialFunc() {
	xxx := map[string]func(){
		"aaa": func() {
			fmt.Println("调用aaa")
		},
		"bbb": func() {
			fmt.Println("调用bbb")
		},
	}
	xxx["aaa"]()
}
func Ff3() {
	incr := func() func() int64 {
		var a int64 = 0
		return func() int64 {
			a++
			return a
		}
	}
	i := incr()
	i2 := incr()
	fmt.Println(i())
	fmt.Println(i2())
	fmt.Println(i())
	fmt.Println(i2())
	fmt.Println(i())
	fmt.Println(i2())
}
