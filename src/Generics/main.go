package main

import "Generics/gen"

func main() {
	//泛型的类型限制，在函数上直接申明该函数支持的多个类型
	//gen.Test1()

	//泛型的类型限制，声明一个interface，包括所有需要支持的类型
	//gen.Test2()

	// 泛型切片,any为泛型自带的一种类型，即interface
	gen.Test3()

	//泛型map
	//gen.Test4()

	//泛型通道
	//gen.Test5()

	//方法约束
	//gen.Test6()

	//类型+方法的约束
	//gen.Test7()

}
