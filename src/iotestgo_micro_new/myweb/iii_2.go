package myweb

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	//"github.com/micro/go-plugins/registry/consul"
	"io/ioutil"
	"log"
	"net/http"
)

func callAPI(addr string, path string, method string) (string, error) {
	req, _ := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient
	res, _ := client.Do(req)
	defer res.Body.Close()
	buf, _ := ioutil.ReadAll(res.Body)
	return string(buf), nil
}
func III3_2() {
	reg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	gets, _ := reg.GetService("xxxyyy")
	//next:=selector.Random(gets)
	next := selector.RoundRobin(gets)
	node, err := next()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(node.Id, node.Address, node.Metadata)
	callres, _ := callAPI(node.Address, "/user", "GET")
	fmt.Println(callres)
}
