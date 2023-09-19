package main

import (
	"5project/myfile"
	"5project/mysock/mynet"
	"5project/mysock/mynet/myrpc"
	"5project/mysock/mynet/myssl"
	"5project/mysock/mynet/myupload"
	"5project/mysock/mynet/template"
	"5project/mysock/myudp/broadcast"
	"5project/mysock/myudp/multi"
)

func main() {
	//写入文件
	myfile.Myfile1()

	//读出文件内容
	//myfile.Myfile2()

	////使用buf_io拷贝文件
	//myfile.Myfile3()

	//使用io_util拷贝文件
	//myfile.Myfile4()

	//扫描文件夹里面的内容
	myfile.Myfile5()

	//文件指针
	//myfile.Myfile6()

	//=========命令行参数============
	//mycmdline.Mycmdline()

	//golang.org/x google的拓展包

	//=========网络io================
	//tcp
	//mytcp.MyTcpServer()
	//mytcp.MyTcpClient()

	//单播
	//onetoone.MyUdpServer()
	//onetoone.MyUdpCli()

	//组播
	multi.MyMultiServer()
	multi.MyMultiCli()

	//广播
	broadcast.MyBCServer()
	broadcast.MyBCCli()

	//=========网络io-http相关================
	mynet.MyHttp1()
	mynet.MyHttp2()
	mynet.MyHttp21()
	mynet.MyHttp22()
	mynet.MyHttp3()
	mynet.MyHttp31()
	mynet.Myhttpcli()

	//模版
	template.TestTemplate()

	//上传
	myupload.MyUpServer()
	myupload.MyUpCli()

	myssl.MySsl()
	myssl.MySslCli()

	//myhttp2.MyHttp2Server()
	//myhttp2.MyHttp2Cli()

	//myws.MyWsServer()

	myrpc.MyRPCServer()
	myrpc.MyRPCCli()
}
