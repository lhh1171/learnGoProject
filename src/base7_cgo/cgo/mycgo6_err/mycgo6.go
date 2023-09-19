package mycgo6_err

//char *cs="hello";
import "C"
import "iotestgo/goandc/cgo/mycgo6_err/hello"

func Mycgo6() {
	//cannot use *_Cvar_cs (type *C.char) as type *hello.C.char in argument to hello.PrintCString2
	hello.PrintCString2(C.cs)
}
