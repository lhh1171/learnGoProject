package mycgo3

//void sayHello(char *s);
import "C"

func Mycgo3(){
	C.sayHello(C.CString("hello world!\n"))
}