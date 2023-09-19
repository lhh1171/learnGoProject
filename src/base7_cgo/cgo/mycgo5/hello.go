package mycgo5

import "C"

import "fmt"

//export sayHello
func sayHello() {
	fmt.Println("hello world!")
}
