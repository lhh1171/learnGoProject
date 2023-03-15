package main

import (
	"3project/myexception"
	"3project/myintf"
	"3project/mystruct"
)

func main() {
	/*========对象成员变量=======*/
	/*指针拼对象*/
	//mystruct.MyStructPtr()

	/*值变量的对象*/
	//mystruct.Mystruct1()

	/*址变量的对象*/
	//mystruct.Mystruct2()

	/*匿名对象*/
	//mystruct.Mystruct3()

	/*值和址的转换自动*/
	//mystruct.Mystruct4()

	/*对象的初始化*/
	//mystruct.Mystruct5()

	/*匿名属性*/
	//mystruct.Mystruct6()

	/*对象构造器*/
	//mystruct.Mystruct7()

	/*tag的使用*/
	//mystruct.Mystruct8()

	/*转类型，stu===>per,子结构体--->父结构体(访问属性) 父结构体--->子结构体(unsafe 通过函数参数隐式转换，内核链表)*/
	//mystruct.Mystruct9()

	/*自定义New*/
	mystruct.Mystruct10()

	/*======对象的成员方法=====*/
	/*值和址的对象调用函数*/
	//mymethod.Mymethod1()

	/*继承的对象调用方法*/
	//mymethod.Mymethod2()

	/*type生成类绑定方法*/
	//mymethod.Mymethod3()

	/*使用new，取出指针，赋值*/
	//mymethod.Mymethod4()

	/*析构函数*/
	//mymethod.Finallizefunc()

	/*========接口================*/
	/*ops当做接口,动态代理*/
	myintf.TestOps()

	/*转类型，用interface
	obj:=接口类型对象.(type)
	obj:=接口类型对象.(目标类型对象)*/
	//myintf.TestPP()

	/*空对象也可以调用函数（特别），
	go语言的interface不影响数据段，只影响代码段*/
	//myintf.Myintf2()

	/*回圈引用,栈溢出*/
	//myintf.TestCircle1()

	/*回圈引用 分包写的 底层也是一个栈溢出 import的时候循环调用init方法*/
	//myintf.TestCircle2()

	/*正常的*/
	//myintf.TestCircle3()

	/*正确解决回圈引用*/
	//myintf.TestCircle4()

	/*exception错误
	panic----throw new exception
	recover----catch(整个函数就是try)*/
	/*捕捉错误err*/
	//myexception.Myexception1()

	/*自定义err*/
	//myexception.Myexception2()

	/*recover&&panic funcb1 funcb21 funcb22里面写的核心*/
	//myexception.Myexception3()

	/*运用deffer打印错误信息*/
	//myexception.Myexception4()

	/*先panic 运用defer最后打印错误信息和recover*/
	myexception.Myexception5()

	/*===========模版模式=========*/
	/*实现方法继承的Override*/
	//template.TestOverride()

	/*两次实现模版模式*/
	//template.TestITemp()
	//template.TestSmsSend()
}
