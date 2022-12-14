package mygoroutine

import (
	"fmt"
	"strconv"
)

func Myselect1() {
	c1 := make(chan int)
	go func() {
		for v := range c1 {
			fmt.Println(v)
		}
	}()

	for i := 0; i < 5; i++ {
		//有点订阅发布的意思，如果之前没有订阅就会死锁
		select {
		case c1 <- 0:
		case c1 <- 1:
		}
	}
}
func Myselect2() {
	c11 := make(chan int)
	c12 := make(chan int)
	go func() {
		c11 <- 1
	}()
	go func() {
		c12 <- 2
	}()
	for i := 0; i < 2; i++ {
		select {
		case j := <-c11:
			fmt.Println("c11: " + strconv.Itoa(j))
		case j := <-c12:
			fmt.Println("c12: " + strconv.Itoa(j))
		}
	}
}
func Myselect3() {
	a := make(chan string)
	select {
	case str := <-a:
		fmt.Println(str)
	default:
		fmt.Println("非阻塞") //否则就是死锁错误
	}

	select {
	case a <- "xxx":
		fmt.Println("发送了xxx")
	default:
		fmt.Println("非阻塞") //否则就是死锁错误
	}
}

