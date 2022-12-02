package myweb

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	myhttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
	"iotestgo_micro_new/modules"
)

func callAPI3(s selector.Selector) {
	mycli := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	req := mycli.NewRequest("xxxyyy", "/user",
		modules.IV4Req{Id: "123"})
	var resp modules.IV4Resp
	mycli.Call(context.Background(), req, &resp)
	fmt.Println(resp.GetData())
}
func IV4_2() {
	reg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	mySelector := selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.Random),
	)
	callAPI3(mySelector)
}
