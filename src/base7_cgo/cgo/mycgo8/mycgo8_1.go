package mycgo8

/*
#include<stdio.h>
char a = -127;
signed char a2 = 254;
unsigned char a3 = 254;
short a4 = -32768;
short unsigned int  a5 = 65535;
int  a6 = -2147483648;
unsigned int  a7 = 4294967295;
long int  a8 = 2147483647;
long unsigned int  a9 = 4294967295;
long long int  a10 = -9223372036854776001;
long long unsigned int  a11 = -9223372036854776001;
*/
import "C"

import (
	"fmt"
	"reflect"
)

func Mycgo8_1() {
	println("==================char —>byte=====================================")
	b := byte(C.a)
	fmt.Println(C.a)
	fmt.Println("类型为:", reflect.TypeOf(C.a))
	fmt.Println(b)
	fmt.Println("类型为:", reflect.TypeOf(b))

	println("===================")
	b2 := int8(C.a)
	fmt.Println(C.a)
	fmt.Println("类型为:", reflect.TypeOf(C.a))
	fmt.Println(b2)
	fmt.Println("类型为:", reflect.TypeOf(b2))
	//=====================================================================
	println("==================signed char—>int8=====================================")

	b3 := byte(C.a2)
	fmt.Println(C.a2)
	fmt.Println("类型为:", reflect.TypeOf(C.a2))
	fmt.Println(b3)
	fmt.Println("类型为:", reflect.TypeOf(b3))
	//=====================================================================
	println("==================unsigned char---->uint8=====================================")
	b4 := byte(C.a3)
	fmt.Println(C.a3)
	fmt.Println("类型为:", reflect.TypeOf(C.a3))
	fmt.Println(b4)
	fmt.Println("类型为:", reflect.TypeOf(b4))
	//=====================================================================
	println("==================short int ----->int16=====================================")
	b5 := int16(C.a4)
	fmt.Println(C.a4)
	fmt.Println("类型为:", reflect.TypeOf(C.a4))
	fmt.Println(b5)
	fmt.Println("类型为:", reflect.TypeOf(b5))
	//=====================================================================
	println("==================short unsigned int ---->uint16=====================================")
	b6 := uint16(C.a5)
	fmt.Println(C.a5)
	fmt.Println("类型为:", reflect.TypeOf(C.a5))
	fmt.Println(b6)
	fmt.Println("类型为:", reflect.TypeOf(b6))
	//=====================================================================
	println("==================int —> int=====================================")
	b7 := int(C.a6)
	fmt.Println(C.a6)
	fmt.Println("类型为:", reflect.TypeOf(C.a6))
	fmt.Println(b7)
	fmt.Println("类型为:", reflect.TypeOf(b7))
	//=====================================================================
	println("==================unsigned int ----->uint32=====================================")
	b8 := uint32(C.a7)
	fmt.Println(C.a7)
	fmt.Println("类型为:", reflect.TypeOf(C.a7))
	fmt.Println(b8)
	fmt.Println("类型为:", reflect.TypeOf(b8))
	//=====================================================================
	println("================== long int —>int32 or int64=====================================")
	b9 := uint32(C.a8)
	c := uint64(C.a8)
	fmt.Println(C.a8)
	fmt.Println("类型为:", reflect.TypeOf(C.a8))
	fmt.Println(b9)
	fmt.Println("类型为:", reflect.TypeOf(b9))
	fmt.Println(c)
	fmt.Println("类型为:", reflect.TypeOf(c))
	//=====================================================================
	println("================== long unsigned int----->uint32 or uint64=====================================")
	b10 := uint32(C.a9)
	c2 := uint64(C.a9)
	fmt.Println(C.a9)
	fmt.Println("类型为:", reflect.TypeOf(C.a9))
	fmt.Println(b10)
	fmt.Println("类型为:", reflect.TypeOf(b10))
	fmt.Println(c2)
	fmt.Println("类型为:", reflect.TypeOf(c2))
	//=====================================================================
	println("================== long long int ----->int64=====================================")
	b11 := int32(C.a10)
	c3 := int64(C.a10)
	fmt.Println(C.a10)
	fmt.Println("类型为:", reflect.TypeOf(C.a10))
	fmt.Println(b11)
	fmt.Println("类型为:", reflect.TypeOf(b11))
	fmt.Println(c3)
	fmt.Println("类型为:", reflect.TypeOf(c3))
	//=====================================================================
	println("================== long long unsigned int —> uint64=====================================")
	b12 := uint32(C.a11)
	c4 := uint64(C.a11)
	fmt.Println(C.a11)
	fmt.Println("类型为:", reflect.TypeOf(C.a11))
	fmt.Println(b12)
	fmt.Println("类型为:", reflect.TypeOf(b12))
	fmt.Println(c4)
	fmt.Println("类型为:", reflect.TypeOf(c4))
}
