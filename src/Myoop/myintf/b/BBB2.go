package b

import (
	"Myoop/myintf/c"
	"fmt"
)

type BBB struct {
	Name string
}

func (a BBB) GetName() string {
	return a.Name
}

func (*BBB) DisBBB(obj interface{}) {
	fmt.Println(obj.(c.CCC).GetName())
}
