package abc

import (
	"fmt"
	"strconv"
)

type Person struct {
	Id   int    //8
	Name string //16
}

func (p Person) setPerson(i int, s string) {
	p.Id = i
	p.Name = s
}
func (p Person) Dis() {
	fmt.Println(strconv.Itoa(p.Id) + "\t" + p.Name)
}
func myOop() {
	p := Person{1, "oooo"}
	Person.setPerson(p, 99, "xxx")
	fmt.Println(p)
}
