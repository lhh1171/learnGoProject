package a

import (
	"Myoop/myintf/c"
	"fmt"
)

type AAA struct {
	Name string
}

func (a AAA) GetName() string {
	return a.Name
}

//func (*AAA) DisAAA(obj *b.BBB) {
//	fmt.Println(obj.Name)
//}

func (*AAA) DisAAA(obj interface{}) {
	fmt.Println(obj.(c.CCC).GetName())
}
