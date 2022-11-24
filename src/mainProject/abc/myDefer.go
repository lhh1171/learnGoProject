package abc

import "fmt"

func MyDefer() {
	res := xxx()
	fmt.Println(res)
}
func MyDefer2() int {
	for i := 0; i < 2; i++ {
		defer fmt.Println(i)
	}
	for i := 0; i < 2; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	return 1
}
func xxx() int {
	defer a1()
	defer a2()
	defer a3()
	fmt.Println("xxxxxx")
	return 1
}
func a1() {
	fmt.Println("11111")
}
func a2() {
	fmt.Println("222222")
}
func a3() {
	fmt.Println("3333333")
}
