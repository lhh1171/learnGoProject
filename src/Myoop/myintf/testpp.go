package myintf

import "fmt"

type Idao interface {
	Save(int) int
}

type Person struct {
	id int
}

func (p Person) Save(a int) int {
	return 99
}

type Car struct {
	name string
}

func (c Car) Save(a int) int {
	return 100
}

func xxx(dao Idao) {
	fmt.Println(dao.Save(88))
	//转换类型方式1
	//c, err := dao.(Car)
	//if err {
	//	fmt.Println(c.name)
	//}
	//
	//p, err2 := dao.(Person)
	//if err2 {
	//	fmt.Println(p.id)
	//}
	//转换类型方式2
	switch d := dao.(type) {
	case Person:
		fmt.Println("person", d.id)
	case Car:
		fmt.Println("car", d.name)
	default:
		fmt.Println("default")
	}
}

func TestPP() {
	xxx(Person{id: 9090})
}
