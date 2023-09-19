package main

import (
	"9project/mygoroutine"
	"9project/mytime"
)

func main() {
	/*========chan的使用=========*/
	//简单使用
	mygoroutine.TestSimple1()

	//只读只写
	//mygoroutine.TestSimple2()

	//err捕捉错误
	//mygoroutine.Mygoroutine1()

	//安排chan对列大小
	//mygoroutine.Mygoroutine2()

	//只读Receive只写Send
	//mygoroutine.Mygoroutine3()

	//close chan 可以解锁
	//mygoroutine.Mygoroutine4()

	//select 有点订阅发布的意思，如果之前没有订阅就会死锁
	//mygoroutine.Myselect1()

	//无阻塞读取
	//mygoroutine.Myselect2()
	//mygoroutine.Myselect3()

	/*===========myruntime运行时设置================*/
	//设置多核处理
	//myruntime.Myruntime()

	//获取当前代码的源代码地址，可以跳过多少行
	//myruntime.TestFilename()

	/*================sync同步实现===================*/
	//原子量
	//mysync.MyAtomic()

	//sync.Mutex底层也是原子量
	//mysync.MyMutx()

	//读写锁,sync.RWMutex,有两个信号量
	//mysync.MyRWLock()

	//栅栏锁
	//mysync.MyCountDownLunch()

	//sync.Once只执行一次的锁
	//mysync.MyOnce()

	//defer封装，当栈溢出时候,函数执行
	//mysync.MyOnce2()

	//sync.Pool,非阻塞，对象池， 线程池，警用了GC
	//mysync.MyPooltest()

	/*=================time================*/
	//After了几秒之后干啥
	//mytime.MyTimeout()

	//timer1 := time.NewTimer timer.c是一个chan
	mytime.MyTimer()

	//每隔500毫秒打印一次
	//mytime.MyTicker()

}
