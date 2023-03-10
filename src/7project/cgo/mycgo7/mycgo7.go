package mycgo7

/*
#cgo linux CFLAGS: -DCGO_OS_LINUX=1
#cgo !linux CFLAGS: -DCGO_OS_WIN=1

#if defined(CGO_OS_LINUX)
	char *os="linux";
#elif defined(CGO_OS_WIN)
	char *os="win";
#else
	char *os="hehe"
#endif
 */
import "C"

func Mycgo7(){
	println(C.GoString(C.os))
}
