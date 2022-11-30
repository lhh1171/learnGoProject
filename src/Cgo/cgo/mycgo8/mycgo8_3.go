package mycgo8

/*
#include<stdio.h>
int  *a14 = 65535;
void *p =&a14;
void printpoint(){
	printf("point:%p\n",p);
}
 */
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)
func Mycgo8_3() {
	println("==================void * ---->*=====================================")
	b := unsafe.Pointer(C.p)
	fmt.Println(C.p)
	fmt.Println("类型为:",reflect.TypeOf(C.a14))
	fmt.Println(b)
	fmt.Println("类型为:",reflect.TypeOf(b))
	fmt.Println(C.printpoint())

}