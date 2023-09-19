package cgo

/*
#include<errno.h>

static int div(int a,int b){
	if(b==0){
		errno=EINVAL;
		return 0;
	}
	return a/b;
}

static void div2(int a,int b){
	if(b==0){
		errno=EINVAL;
	}
}
*/
import "C"
import "fmt"

func Mycgo9() {
	v0, err0 := C.div(1, 2)
	fmt.Println(v0, err0)

	v0, err0 = C.div(1, 0)
	fmt.Println(v0, err0)

	_, err0 = C.div2(1, 0)
	fmt.Println(err0)
}
