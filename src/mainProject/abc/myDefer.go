package abc

import "fmt"

func MyDefer() {
	res := xxx()
	fmt.Println(res)
}
func MyDefer2() int {
	for i := 0; i < 2; i++ {
		//打印1 0
		defer fmt.Println(i)
	}
	for i := 0; i < 2; i++ {
		//打印2 2
		/*闭包了，i属于外部引用*/
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
	return 111
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
