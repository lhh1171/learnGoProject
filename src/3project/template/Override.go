package template

import (
	"fmt"
)

type ITmp1 interface {
	load()
}

// Tmp 基类
type tmp1 struct {
	ITmp
}

// Load 基类方法
func (t *tmp) load() {
	fmt.Println("--------default---------")
	t.load()
}

// Aaax 实现类
type Aaa struct {
	*tmp
}

func NewAaa2() *Aaa {
	aa := &Aaa{}
	aa.tmp = &tmp{ITmp: aa}
	return aa
}

// Load 实现方法
func (a *Aaa) load() error {
	fmt.Println("Aaa Load..........")
	return nil
}

func TestOverride() {
	aa2 := NewAaa2()
	aa2.load()
}
