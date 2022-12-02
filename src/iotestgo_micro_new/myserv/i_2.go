package myserv

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	//_ "github.com/micro/go-plugins/registry/consul/v2"
	"golang.org/x/net/context"
	"iotestgo_micro_new/proto"
)

func I1_2() {
	service := micro.NewService(
		micro.Name("client.service"),
	)

	service.Init()

	greeter := proto.NewGreeterService("go.micro.api.greeter", service.Client())

	rsp, err := greeter.HelloTest(context.TODO(), &proto.SayRequest{
		Name: "Hello test",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.GetGreeting)
}
