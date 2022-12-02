package myweb

import (
	"github.com/micro/go-micro/web"
	"net/http"
)

func I1() {
	server:=web.NewService(web.Address("0.0.0.0:9999"))
	server.HandleFunc("/",func(resp http.ResponseWriter,req *http.Request){
		resp.Write([]byte("hello world"))
	})
	server.Run()
}
