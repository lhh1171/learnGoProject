package cgo

/*
#include<stdio.h>

static void sayHello(char *s){
	puts(s);
}
*/
import "C"

func Mycgo2() {
	C.sayHello(C.CString("hello world!\n"))
}
