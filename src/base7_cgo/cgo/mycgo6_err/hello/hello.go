package hello

//#include<stdio.h>
import "C"

func PrintCString2(cs *C.char) {
	C.puts(cs)
}
