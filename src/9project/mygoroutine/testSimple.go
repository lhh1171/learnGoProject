package mygoroutine

import (
	"fmt"
	"time"
	"unsafe"
)

//	func run(c *chan int8) {
//		var a int8 = 99
//		time.Sleep(1 * time.Second)
//		fmt.Println("=====================")
//		//对chan赋值
//		//下面写东西没有用了
//		*c <- a
//		//time.Sleep(1 * time.Second)
//		fmt.Println("=====================")
//	}
func run(c *chan int8) {
	var a int8 = 99
	time.Sleep(1 * time.Second)
	//产生了一条数据
	*c <- a
}
func run2(c *chan int8) {
	//消费了两次
	d := <-*c
	//死锁了
	d = <-*c
	fmt.Println(d)
}

func TestSimple1() {
	//返回的是一个对象（句柄），对象里有一个指针
	//chan是为了让协程之间通信
	c := make(chan int8)
	go run(&c)
	go run2(&c)
	time.Sleep(8 * time.Second)
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println("xxxxx")
}

// 单向传递，只写
func run3(c chan<- int8) {
	var a int8 = 99
	time.Sleep(1 * time.Second)
	//产生了一条数据
	c <- a

}

// 单向传递，只读
func run4(c <-chan int8) {
	d := <-c
	fmt.Println(d)
}
func TestSimple2() {
	//返回的是一个对象（句柄），对象里有一个指针
	//chan是为了让协程之间通信
	c := make(chan int8)
	//run3 run4指的是mailBox的两块地方
	go run3(c)
	go run4(c)
	time.Sleep(3 * time.Second)
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println("xxxxx")
}
