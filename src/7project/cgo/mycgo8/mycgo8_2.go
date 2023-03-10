package mycgo8

/*
#include<stdio.h>
float a12 = -3.14;
double a13 = -3.14159267333333333333333;
 */
import "C"

import (
	"fmt"
	"reflect"
)
func Mycgo8_2() {
	println("==================float---->float32=====================================")
	b := float32(C.a12)
	fmt.Println(C.a12)
	fmt.Println("类型为:",reflect.TypeOf(C.a12))
	fmt.Println(b)
	fmt.Println("类型为:",reflect.TypeOf(b))
	//=====================================================================
	println("==================double—>float64=====================================")
	b2 := float32(C.a13)
	c := float64(C.a13)
	fmt.Println(C.a13)
	fmt.Println("类型为:",reflect.TypeOf(C.a13))
	fmt.Println(b2)
	fmt.Println("类型为:",reflect.TypeOf(b2))
	fmt.Println(c)
	fmt.Println("类型为:",reflect.TypeOf(c))
}