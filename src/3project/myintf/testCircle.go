package myintf

import (
	"fmt"
	"iotestgo/myoop/myintf/a"
	"iotestgo/myoop/myintf/b"
)

/*在任何语言写不了*/
//type Aaa struct {
//	b Bbb
//}
//
//type Bbb struct {
//	a Aaa
//}

/*=====================*/

//type Aaa struct {
//	b *Bbb
//}
//
//func (a Aaa) aaa() {
//	a.b.bbb()
//}
//
//type Bbb struct {
//	a *Aaa
//}
//
//func (b Bbb) bbb() {
//	b.a.aaa()
//}
//
//// TestCircle1
/////*栈溢出*/
//func TestCircle1() {
//	a:=Aaa{}
//	b:=Bbb{}
//	a.b=&b
//	b.a=&a
//	a.aaa()
//}

// TestCircle2 /*编译报错
//package command-line-arguments
//	imports Myoop/myintf
//	imports Myoop/myintf/a

//	imports Myoop/myintf/b
//	imports Myoop/myintf/a: import cycle not allowed*/
//func TestCircle2() {
//	a := a2.Aaa{}
//	fmt.Println(a)
//	b := b2.Bbb{}
//	fmt.Println(b)
//}

type Aaaa struct {
	b *Bbbb
}

type Bbbb struct {
	a *Aaaa
}

func TestCircle3() {
	a := Aaaa{}
	fmt.Println(a)
	b := Bbbb{}
	fmt.Println(b)
}

func TestCircle4() {
	obj := a.AAA{"AAA"}
	obj1 := b.BBB{"BBB"}
	obj.DisAAA(&obj1)
	obj1.DisBBB(&obj)
}
