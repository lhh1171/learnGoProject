package mycgo8

/*
#include<stdio.h>
enum C {
    ONE,
    TWO,
};
*/
import "C"
import "fmt"

func Mycgo8_7() {
	println("==================枚举类型=====================================")
	//枚举类型,通过C.enum_xxx访问C语言当中定义的enum xxx结构体类型
	var c C.enum_C = C.TWO
	fmt.Println(c)
	fmt.Println(C.ONE)
}
