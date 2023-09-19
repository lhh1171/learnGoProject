package mycgo8

/*
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static char* cat(char* str1, char* str2) {
	static char buf[256];
	strcpy(buf, str1);
	strcat(buf, str2);

	return buf;
}

char *s = "Hello";
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func Mycgo8_8() {
	println("==================字符串=====================================")
	str1, str2 := "hello", " world"
	// golang string -> c string
	cstr1, cstr2 := C.CString(str1), C.CString(str2)
	defer C.free(unsafe.Pointer(cstr1)) // 该内存不归go管理
	defer C.free(unsafe.Pointer(cstr2))
	cstr3 := C.cat(cstr1, cstr2)
	// c string -> golang string
	str3 := C.GoString(cstr3)
	fmt.Println(str3) // "hello world"

	println("==============")
	var s0 string
	var s0Hdr = (*reflect.StringHeader)(unsafe.Pointer(&s0))
	s0Hdr.Data = uintptr(unsafe.Pointer(C.s))
	s0Hdr.Len = int(C.strlen(C.s))

	println(s0)

	// 通过切片语法转换
	sLen := int(C.strlen(C.s))
	s1 := string((*[16]byte)(unsafe.Pointer(C.s))[:sLen:sLen])
	println(s1)
}
