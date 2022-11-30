package mycgo5

//#include"hello.h"
import "C"

func Mycgo5(){
	C.sayHello(C.CString("hello world!\n"))
}
