package main

import (
	"firstProject1/ccc"
	. "firstProject1/ccc"
	Tu "firstProject1/ccc"
	"fmt"
)

// init方法在import的时候就会被调用
// 导包可以起别名
func main() {
	//--------------------包管理------------------------
	ccc.Aaa()
	Aaa()
	Tu.Aaa()
	fmt.Println("hello world!")
	//----------------------数据类型------------------------
	////没有val
	////整形
	//var x byte = 1
	////byte就相当于uint8
	//fmt.Println(x)
	//var a int8 = 1
	//fmt.Println(a)
	//var b uint8 = 1
	//b = 90
	//fmt.Println(b)
	///*普通变量在栈里，特殊数据类型（指针和类类型），没有函数返回的话都是放在栈上，
	//只要出现逃逸（从你的栈里跑到其他的栈里）就会放在堆上，由gc协程回收*/
	///*malloc(linux默认)，JeMalloc,TCMalloc*/
	////long在32位机里32位，64位机器里64位
	////同理int 16 32
	//var c int = 2 //自动识别，32位机16,64为32
	//fmt.Println(c)
	//var d uintptr = 0 //void *
	//fmt.Println(d)
	////整数     空      指针
	////c:     void*
	////uint_ptr------Unsafe.Point(需要手动转换）
	////&c不能取其他类型
	//var ptr *int
	//ptr = &c
	//fmt.Println("&C", ptr)
	//fmt.Println(*ptr)
	///*指针操作*/
	///*uintptr 是int 与unsafe.Pointer之前的桥梁*/
	///*61--a,62--b*/
	//var e int16 = 0x6162
	//var ePtr *int16 = &e
	///*0xc000104cba*/
	////只能看见b,a
	//fmt.Printf("0x%x\n", *ePtr)
	////指针取var
	//var l = int64(uintptr(unsafe.Pointer(ePtr)))
	//fmt.Printf("sssssssssss0x%x\n", l)
	////var取指针
	//var l1 *uint8 = (*uint8)(unsafe.Pointer(uintptr(l)))
	//fmt.Println(*l1)
	//fmt.Printf("l10x%x\n", *l1)
	//l++
	//var l2 *uint8 = (*uint8)(unsafe.Pointer(uintptr(l)))
	//fmt.Println(*l2)
	//fmt.Printf("l20x%x\n", *l2)
	////浮点
	//var f float32 = 1.1
	//fmt.Println(f)
	//type
	/*	type size int
		var xx size = 1
		var yy int = 1
		//yy= int(xx)
		type mysize = int
		var zz mysize = 1
		fmt.Println(reflect.TypeOf(xx).Kind()) //判断根类型
		fmt.Println(reflect.TypeOf(xx), reflect.TypeOf(yy), reflect.TypeOf(zz))*/
	//布尔
	/*	var bl bool = true
		fmt.Println(bl)*/
	//string
	/*	var str string = "abc"
			var str2 string = "中国"
			var str3 string = `asdf
		aasdfa
		asdfas`
			fmt.Println(str3)
			fmt.Println(len(str))
			fmt.Println(len(str2))
			fmt.Println(utf8.RuneCountInString(str))
			fmt.Println(utf8.RuneCountInString(str2))
			var str4 string = "zgg"
			var sPtr = (*reflect.StringHeader)(unsafe.Pointer(&str4))
			fmt.Println(sPtr.Len)
			fmt.Println(sPtr.Data)
			var cc *uint8 = (*uint8)(unsafe.Pointer(sPtr.Data))
			fmt.Printf("%c\n", *cc)
			fmt.Println(unsafe.Sizeof(sPtr))*/
	//运算符
	/*	var aa byte = 0xd0 //1101 0000
		var bb byte = 0x10 //0001 0000
		//1100 0000特殊位运算
		var ab = aa & ^bb
		//异或相同为0不同为1
		//得0x80 1000 0000
		fmt.Printf("0x%x\n", ab)*/
	//流程控制 if switch for

	/*	if la := 1; la == 1 {
			fmt.Println("111")
		} else if la == 2 {

		} else {

		}*/
	//la在if 那个栈里没有逃逸 所以找不到
	//la=2

	/*=================*/

	//if ab==2{
	//	fmt.Println("ab==2")
	//}else if ab==3 {
	//	fmt.Println("ab==3")
	//}else {
	//	fmt.Println("ab==else")
	//}

	/*============*/
	//fallthrough表示可以往下走
	/*	switch sw := 1; sw {
			case 1:
				fmt.Println("111")
			case 2:
				fmt.Println("222")
				fallthrough
			default:
				fmt.Println("defalut")
			}
			switch sw := "11"; sw {
			case "11":
				fmt.Println("111")
			case "22":
				fmt.Println("222")
				fallthrough
			default:
				fmt.Println("defalut")
			}
		xxx:
			for {
				for fa := 1; fa < 9; fa++ {
					fmt.Println(fa)
					//跳出指定循环 continue也可以
					break xxx
				}
			}*/
	//数组
	/*	var arr1 = [5]int{1, 2, 3, 4, 5}
		for i, v := range arr1 {
			fmt.Println(i, v)
		}
		var id, name, age = 1, "aaa", 22
		fmt.Println(id, name, age)*/
	//交换俩数
	/*	y1 := 1
		y2 := 2
		y1, y2 = y2, y1
		fmt.Println(y1, y2)*/
	//整数<===>字符串
	/*	var sstr = "123"
		//myint, _ := strconv.Atoi(sstr)
		myint, _ := strconv.ParseInt(sstr, 0, 10) //10代表十进制
		fmt.Println(myint + 1)
		var vv = 123
		fmt.Println(strconv.Itoa(vv) + "yu")*/
	//集合
	//数组,内存连续
	//...代表自动推倒，也可以写长度
	//arr是连续内存
	/*	arr := [...]uint8{1, 2, 3, 4}
		ii := uint(uintptr(unsafe.Pointer(&arr)))
		ii += 3
		aptr := (*uint8)(unsafe.Pointer(uintptr(ii)))
		fmt.Println(*aptr)
		fmt.Println(arr, reflect.TypeOf(arr))*/
	//slice
	//没有写长度就是sliceHeader和stringHeader差不多，里面
	/*	sl := []uint8{1, 2, 3, 4, 5}
		fmt.Println(sl, reflect.TypeOf(sl))
		slptr := (*reflect.SliceHeader)(unsafe.Pointer(&sl))
		fmt.Println(slptr.Len, "\t", slptr.Cap, "\t", slptr.Data)*/
	//slptr.Data++
	/*	uptr := (*uint8)(unsafe.Pointer(uintptr(slptr.Data)))
		fmt.Println(*uptr)
		sl = append(sl, 6, 7)
		fmt.Println(sl)
		//map
		var m map[string]int = map[string]int{"aaa": 111, "bbb": 222}
		fmt.Println(m)
		//getForMap
		fmt.Println(m["aaa"])*/
	//set
	/*	m["fdsa"] = 222
		for k, v := range m {
			fmt.Println(k, v)
		}
		fmt.Println(unsafe.Sizeof(m))*/
}
