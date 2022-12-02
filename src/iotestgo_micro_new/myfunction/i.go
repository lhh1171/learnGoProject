package myfunction

import (
	"context"
	"iotestgo_micro_new/proto"

	"github.com/micro/go-micro/v2"
)

type Say struct{}

func (g *Say) HelloTest(ctx context.Context, req *proto.SayRequest, rsp *proto.SayResponse) error {
	rsp.Greeting = "你好呀！ " + req.Name
	return nil
}

func I1() {
	fnc := micro.NewFunction(
		micro.Name("greeter"),
	)

	fnc.Init()

	fnc.Handle(new(Say))

	fnc.Run()
}

