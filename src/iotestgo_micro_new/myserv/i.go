package myserv

import (
	"context"
	"fmt"
	"iotestgo_micro_new/proto"
	"github.com/micro/go-micro/v2"
)

type Greeter struct {
}

func (g *Greeter) HelloTest(ctx context.Context, req *proto.SayRequest, resp *proto.SayResponse) error {
	fmt.Println("recived data")
	resp.Greeting = req.Name + "===="
	return nil
}

func I1() {
	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
		)
	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

