package mycgo4

//void sayHello(char *s);
import "C"

func Mycgo4() {
	C.sayHello(C.CString("hello world!\n"))
}
