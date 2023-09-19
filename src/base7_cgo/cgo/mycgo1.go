package cgo

//#include<stdio.h>
import "C"

func Mycgo1() {
	C.puts(C.CString("hello world!\n"))
}
