package myintf

import (
	"3project/myintf/a"
	"3project/myintf/b"
	"fmt"
)

/*在任何语言写不了,栈溢出，Invalid recursive type 'Bbb' Bbb → Aaa → Bbb,解析出错*/
/*type Aaa struct {
	b Bbb
}

type Bbb struct {
	a Aaa
}*/
/*=========回圈引用============*/

/*type Aaa struct {
	b *Bbb
}

func (a Aaa) aaa() {
	a.b.bbb()
}

type Bbb struct {
	a *Aaa
}

func (b Bbb) bbb() {
	b.a.aaa()
}*/

// TestCircle1
// /*栈溢出*/
/*
func TestCircle1() {
	a := Aaa{}
	b := Bbb{}
	a.b = &b
	b.a = &a
	a.aaa()
}*/

// TestCircle2 /*编译报错
/*package command-line-arguments
imports Myoop/myintf
imports Myoop/myintf/a

imports Myoop/myintf/b
imports Myoop/myintf/a: import cycle not allowed*/

/*func TestCircle2() {
	a := a.Aaa{}
	fmt.Println(a)
	b := b.Bbb{}
	fmt.Println(b)
}*/

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
