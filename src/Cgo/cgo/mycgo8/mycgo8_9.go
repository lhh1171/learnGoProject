package mycgo8

/*
#include <stdint.h>

static void fill_255(char* buf, int32_t len) {
	int32_t i;
	for (i = 0; i < len; i++) {
		buf[i] = 255;
	}
}

char arr[10]={1,2,3,4,5,6,7,8,9,10};
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func Mycgo8_9() {
	println("==================切片=====================================")
	b := make([]byte, 5)
	fmt.Println(b) // [0 0 0 0 0]
	C.fill_255((*C.char)(unsafe.Pointer(&b[0])), C.int32_t(len(b)))
	fmt.Println(b) // [255 255 255 255 255]

	println("==============")
	// 通过 reflect.SliceHeader 转换
	var arr0 []byte
	var arr0Hdr = (*reflect.SliceHeader)(unsafe.Pointer(&arr0))
	arr0Hdr.Data = uintptr(unsafe.Pointer(&C.arr[0]))
	arr0Hdr.Len = 10
	arr0Hdr.Cap = 10
	for v:=range arr0{
		print(v," ")
	}
	println()
	// 通过切片语法转换
	arr1 := (*[24]byte)(unsafe.Pointer(&C.arr[0]))[:10:10]
	for v:=range arr1{
		print(v," ")
	}
}