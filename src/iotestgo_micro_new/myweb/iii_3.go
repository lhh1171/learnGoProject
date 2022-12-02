package myweb

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	myhttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
)

func callAPI2(s selector.Selector) {
	mycli := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	req := mycli.NewRequest("xxxyyy", "/user", map[string]string{"id": "www", "name": "ppp"})
	var resp map[string]interface{}
	mycli.Call(context.Background(), req, &resp)
	fmt.Println(resp["www"])
}
func III3_3() {
	reg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	mySelector := selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.Random),
	)
	callAPI2(mySelector)
}
