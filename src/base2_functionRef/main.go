package main

import (
	"2project/myref"
)

func main() {
	/*函数声明和调用*/
	//abc.Ff1()
	//abc.Ff2()
	/*偏函数*/
	//abc.PartialFunc()
	/*闭包*/
	//abc.Ff3()
	/*defer函数  压栈绑定的是外部函数*/
	//abc.MyDefer()
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

	//Const
	//abc.TestConst()

	/*=================反射reflect====================*/
	//TypeOf获取字段属性以及字段相关的东西,ValueOf获取字段的值以及值相关的东西
	//myref.Myref1()

	//Field获取成员变量的name和type,Method获取方法的name和type
	//myref.Myref2()

	//指针通过elem获取内容，通过method获取方法,通过elem仅能获取非指针对象的引用方法
	//myref.Myref3()

	//获取成员FieldByName，FieldByIndex，FieldByIndex，MethodByName
	//myref.Myref4()

	//反射基础数据类型,elem的set和canSet(有点像右值)，Call调用函数
	//myref.Myref5()

	//reflect.TypeOf  返回Type对象 查看里面的属性
	//myref.Myref6()

	//运用反射构建map量表，classForName
	myref.Myref7()

}
