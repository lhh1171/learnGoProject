package main

import (
	"6project/myapi"
)

func main() {
	//=========基本api================
	/*对数组的基本操作,类似于文件偏移量*/
	//bytes.BytesBase()
	//bytes.BytesReader()
	//bytes.BytesBuffer()
	/*正则*/
	//myapi.MyReg()
	/*字符串*/
	//myapi.Mystr()
	/*时间戳*/
	//myapi.Mytime()

	//myapi.MyBase64()
	//myapi.MyJson()

	/*输入输出*/
	//myapi.MyLineFilter()

	//myapi.MyRandom()
	//myapi.MySha1()
	myapi.MyString2()
	/*解析url*/
	myapi.MyUrlParser()
}
