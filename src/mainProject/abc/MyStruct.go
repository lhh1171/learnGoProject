package abc

import (
	"fmt"
	"reflect"
	"unsafe"
)

func MyStruct() {
	type Person struct {
		Id   int    //8
		Name string //16
	}
	P := Person{1, "xxx"}
	fmt.Println(P.Id, P.Name)
	fmt.Println(unsafe.Sizeof(P))
}
func MyStruct1() {
	type Person struct {
		int    //8
		string //16
	}
	P := Person{1, "xxx"}
	fmt.Println(P.int, P.string)
}
func MyStruct2() {
	type Person struct {
		Id   int    //8
		Name string //16
	}
	var arr [3]int64 = [3]int64{}
	arr[0] = 99
	arr[2] = 3
	str := "abc"
	ptr := (*reflect.StringHeader)(unsafe.Pointer(&str))
	arr[1] = int64(ptr.Data)
	p := (*Person)(unsafe.Pointer(&arr))
	fmt.Println(p.Id, "\t", p.Name)
}
func MyStruct3() {
	type Person struct {
		Id   int    //8
		Name string //16
	}
	type stu struct {
		Person
		age int
	}
	s := stu{
		age: 1,
		Person: struct {
			Id   int
			Name string
		}{},
	}
	s.age = 22
	s.Id = 1
	s.Name = "aaa"
	fmt.Println(s.age, s.Name, s.Id)
}
