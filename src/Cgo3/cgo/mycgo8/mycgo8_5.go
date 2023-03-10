package mycgo8

/*
#include<stdio.h>
struct A {
    int i;
    float f;
};

struct A2 {
    int type;
    char chan;
};

struct A3 {
    int size:10; // 位字段无法访问
    float arr[]; // 零长的数组也无法访问
};
 */
import "C"
import "fmt"

func Mycgo8_5() {
	println("==================结构体=====================================")
	var a C.struct_A
	fmt.Println(a.i)
	println("============")
	//结构体当中出现了Go语言的关键字,通过下划线的方式进行访问
	//注意：如果有两个成员,一个是以Go语言关键字命名,另外一个刚好是以下划线和Go语言关键字命名,那么以Go语言关键字命名的成员将无法访问(被屏蔽)
	var a2 C.struct_A2
	fmt.Println(a2._type)
	fmt.Println(a2._chan)
	println("============")
	//goandc/cgo/mycgo8/mycgo8_5.go:35: a3.size undefined (type C.struct_A3 has no field or method size)
	//goandc/cgo/mycgo8/mycgo8_5.go:36: a3.arr undefined (type C.struct_A3 has no field or method arr)
	//var a3 C.struct_A3
	//fmt.Println(a3.size)
	//fmt.Println(a3.arr)
	//在 C 语言当中, 无法直接访问 Go 语言定义的结构体类型
}