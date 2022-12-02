package main

import "iotestgo_micro_new/myweb"

func main() {
	//http api部分
	myweb.I1()
	//myweb.II2()
	//myweb.III3()
	//myweb.III3_2()
	//myweb.III3_3()
	//myweb.IV4()
	//myweb.IV4_2()
	//grpc部分,从该部分起必须用v2版本
	//go get github.com/micro/micro/v2
	//myserv.I1()
	//myserv.I1_2()
	//broker部分
	//mybroker.I1()
	//mybroker.I1_2()

	//mybroker.II2()
	//mybroker.II2_2()
	//gateway部分

	//gateway--api
	//  /disk/go13pro/bin/micro api --handler=api
	//myapi.I1()
	//curl -H 'head-1: I am a header' "http://localhost:8080/example/call?name=john"
	//curl -H 'Content-Type: application/json' -d '{data:123}' http://localhost:8080/example/foo/bar

	//gateway --rpc
	//  /disk/go13pro/bin/micro api --handler=rpc
	//curl -H 'head-1: I am a header' "http://localhost:8080/example/call?name=john"

	//gateway --event
	//  /disk/go13pro/bin/micro api --handler=event --namespace=go.micro.evt
	//myapi.II2()
	//curl -d '{"message": "Hello, Micro中国"}' http://localhost:8080/user/login

	//gateway --meta
	//myapi.III3()
	//  /disk/go13pro/bin/micro api
	//向 /example POST请求时，会被转到go.micro.api.example的Example.Call方法。
	//curl -H 'Content-Type: application/json' -d '{"name": "john"}' "http://localhost:8080/example"
	//curl -XGET "http://localhost:8080/example?name=john"
	//向 /example POST请求时，会被转到go.micro.api.example的Foo.Bar方法。
	//
	//curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/foo/bar
	//curl -XGET "http://localhost:8080/foo/bar"
	//curl -XDELETE "http://localhost:8080/foo/bar"

	//function
	//myfunction.I1()
	// /disk/go13pro/bin/micro call greeter Say.HelloTest '{"name": "xxxx"}'
}
