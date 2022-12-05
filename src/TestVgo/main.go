package main

import (
	"context"
	"fmt"
	"time"

	proto "TestVgo/proto"
	"github.com/micro/micro/v3/service"
)

func main() {
	// create and initialise a new service
	srv := service.New()

	// create the proto client for helloworld
	client := proto.NewHelloworldService("helloworld", srv.Client())

	// call an endpoint on the service
	rsp, err := client.Call(context.Background(), &proto.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println("Error calling helloworld: ", err)
		return
	}

	// print the response
	fmt.Println("Response: ", rsp.Msg)

	// let's delay the process for exiting for reasons you'll see below
	time.Sleep(time.Second * 5)
}
