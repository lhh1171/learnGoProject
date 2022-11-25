package mystruct

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

type Obj struct {
	id   int
	name string
}

type person struct {
	id   int
	name string
}

func Mystruct1() {

	var p1 = person{1, "aaa"}
	var p2 = person{}
	p2.id = 2
	p2.name = "bbb"
	fmt.Println(p1, p2)
}
func Mystruct2() {
	var p3 = new(person)
	p3.id = 3
	p3.name = "ccc"
	(*p3).id = 3
	(*p3).name = "ccc"
	fmt.Println(p3, *p3)
}
func Mystruct3() {
	var tmpperson = struct {
		id   int
		name string
	}{1, "aaa"}
	fmt.Println(tmpperson)

	fmt.Println("======")

	var tmppersonptr = &struct {
		id   int
		name string
	}{1, "aaa"}
	fmt.Println(tmppersonptr, *tmppersonptr)
}
func Mystruct4() {
	var o1 = Obj{1, "aaa"}
	var o2 = &Obj{2, "bbb"}
	changeObj(o1, o2)
	fmt.Println(o1, o2)
}
func Mystruct5() {
	var car = struct {
		id    int
		name  string
		style struct {
			color int
			model string
		}
	}{1, "aaa", struct {
		color int
		model string
	}{color: 1, model: "xxx"}}
	fmt.Println(car)

	fmt.Println("======")

	car.id = 2
	car.name = "bbb"
	car.style.color = 2
	car.style.model = "yyy"
	fmt.Println(car)
}
func Mystruct6() {
	var cartemp1 = struct {
		int
		string
	}{1, "xxx"}
	fmt.Println(cartemp1)

	fmt.Println("======")

	var cartemp2 = struct {
		int
		string
		//string//使用匿名属性，属性数据类型不能重复，其实质上是把类型当做属性名使用了
		style struct {
			int
			string
		}
	}{1, "xxx", struct {
		int
		string
	}{2, "yyy"}}
	fmt.Println(cartemp2)

	fmt.Println(cartemp2.int, cartemp2.string, cartemp2.style.int, cartemp2.style.string)

}
func Mystruct7() {
	type Person struct {
		id   int
		name string
	}
	type stu struct {
		Person
		age int
	}
	type tea struct {
		Person
		cource string
	}
	var s = stu{
		Person: Person{id: 1, name: "aaa"},
		age:    22,
	}
	var t = tea{}
	t.Person.id = 2
	t.id = 2
	t.Person.name = "bbb"
	t.name = "bbb"
	t.cource = "ccc"

	fmt.Println(s, t)
}
func Mystruct8() {
	type Base struct {
		basename string
	}
	type Derive struct {
		*Base
		id int
	}
	var d = Derive{&Base{basename: "xxx"}, 1}
	fmt.Println(d)

	fmt.Println("========================8=======================")

	//========================================================================

	type Mytag struct {
		X1 string `k:"aaa" v:"bbb" x:"ppp"`
	}
	var mt Mytag
	ct, ok := reflect.TypeOf(mt).FieldByName("X1")
	if ok {
		fmt.Println(ct.Tag.Get("k"), ct.Tag.Get("v"), ct.Tag.Get("x"))
	}
}
func (p person) Dis() {
	fmt.Println(strconv.Itoa(p.id) + "\t" + p.name)
}
func (s Stu) Dis() {
	fmt.Println(strconv.Itoa(s.id) + "\t" + s.name + "\t" + "stu........")
}

type Stu struct {
	person
	age int
}

func Mystruct9() {
	s := Stu{
		person: struct {
			id   int
			name string
		}{
			id:   12,
			name: "aaaa",
		},
		age: 34,
	}
	s.Dis() //stu
	ptr := (*person)(unsafe.Pointer(&s))
	ptr.Dis() //per
}

func (s Stu) New(id int, name string, age int) Stu {
	return Stu{
		person: struct {
			id   int
			name string
		}{
			id:   id,
			name: name,
		},
		age: age,
	}
}
func Mystruct10() {
	ss := Stu.New(Stu{}, 1, "2", 3)
	ss.Dis()
}

func changeObj(o1 Obj, o2 *Obj) {
	//自动类型转换
	o1.id = 3
	o1.name = "ccc"
	o2.id = 4
	o2.name = "ddd"
}
