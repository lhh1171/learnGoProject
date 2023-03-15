package main

func main() {
	//fmt.Println("fsdafa")
	/*执行系统命令*/
	//MyExec()
	MyExec2()

	/*kill -9发送信号量*/
	MySignal()

	/*。我们应该总是使用 Syscall，
	RawSyscall 存在的意义是为那些永远不会阻塞的系统调用准备的，*/
	MySyscall()
	MySyscall2()
}
