package main

import "mainProject/abc"

func main() {
	/*函数声明和调用*/
	//abc.Ff1()
	//abc.Ff2()
	/*偏函数*/
	//abc.PartialFunc()
	/*闭包*/
	//abc.Ff3()
	/*defer函数*/
	abc.MyDefer()
	//abc.MyDefer2()
	//defer
	//abc.MyStruct()
	//abc.MyStruct1()
	//结构体对齐，go哪里缺哪里对齐，java是最后补齐
	/*
		{int int8 string}//8+1+(7)+16
		补了个7
	*/
	//abc.MyStruct2()
	//abc.MyStruct3()
}
