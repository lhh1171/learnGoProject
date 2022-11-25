package myintf

import "fmt"

type Iaa interface {
	iaa()
}
type Ibb interface {
	ibb()
}
type IA struct {
}

func (self IA) iaa() {
	fmt.Println("iaaaaaaaaaaa")
}

type IB struct {
	IA
}

func (self IB) ibb() {
	fmt.Println("ibbbbbbbbbbbbb")
}

// interface{}相当于object
func xxxx(ia interface{}) {
	//一切都是interface，指针也是interface
	//不关心内存结构，关心类型系统
	ib, _ := ia.(IB)
	//ib.IA.iaa() 自动转换
	ib.iaa()
	ib.ibb()
}
func TestAB() {
	ia := IB{}
	xxxx(&ia)
}
