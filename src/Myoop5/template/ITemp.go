package template

import (
	"fmt"
)

type ITmp interface {
	load() error
}

// Tmp 基类
type tmp struct {
	ITmp
}

// Load 基类方法
func (t *tmp) load1() error {
	fmt.Println("--------default---------")
	return t.load()
}

// Aaax 实现类
type Aaax struct {
	*tmp
}

func NewAaa() *Aaax {
	aa := &Aaax{}
	aa.tmp = &tmp{ITmp: aa}
	return aa
}

// Load 实现方法
func (a *Aaax) load() error {
	fmt.Println("Aaa Load..........")
	return nil
}

func TestITemp() {
	aa := NewAaa()
	aa.load1()
}
