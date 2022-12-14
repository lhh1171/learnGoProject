package mymethod

import (
	"fmt"
	"runtime"
	"time"
)

type Data struct {
	Name string
	d    [1024 * 100]byte
}

func close1(d *Data) {
	fmt.Println("被回收了！！", d.Name)
}
func test() {
	var a Data
	a.Name = "xxx"
	fmt.Println("yyy")
	//回收的时候(栈结束的时候),调用close
	runtime.SetFinalizer(&a, close1)
}

func Finallizefunc() {
	for {
		test()
		time.Sleep(time.Millisecond)
	}
}
