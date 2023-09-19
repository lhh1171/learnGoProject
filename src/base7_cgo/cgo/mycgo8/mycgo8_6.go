package mycgo8

/*
#include<stdio.h>
#include<stdint.h>
union B {
    int i;
    float f;
};

union C {
    int8_t i8;
    int64_t i64;
};
*/
import "C"
import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

func Mycgo8_6() {
	println("==================共用体=====================================")
	//对于联合类型,可以通过C.union_xxx来访问C语言中定义的union xxx类型.
	//但是Go语言中并不支持C语言联合类型,它们会被转换为对应大小的字节数组.
	var b C.union_B
	fmt.Printf("%T \n", b) // [4]uint8

	var c C.union_C
	fmt.Printf("%T \n", c) // [8]uint8

	//通过 Go 语言的 encoding/binary 手工解码成员(需要注意大小端的问题)
	var ub C.union_B
	fmt.Println("b.i:", binary.LittleEndian.Uint32(ub[:]))
	fmt.Println("b.f:", binary.LittleEndian.Uint32(ub[:]))

	//使用 unsafe 包强制转换为对应的类型
	fmt.Println("b.i:", *(*C.int)(unsafe.Pointer(&ub)))
	fmt.Println("b.f:", *(*C.float)(unsafe.Pointer(&ub)))
}
