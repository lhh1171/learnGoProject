package myweb

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func IV4() {
	reg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	r := gin.Default()
	type Obj struct {
		Id string `form:"id"`
	}
	type Ret struct {
		Id   string
		Name string
	}
	ret := make([]Ret, 0)
	ret = append(ret, Ret{"111", "aaa"})
	ret = append(ret, Ret{"222", "bbb"})
	r.Handle("POST", "/user", func(c *gin.Context) {
		var o Obj
		c.Bind(&o)
		fmt.Println(o)
		c.JSON(200, gin.H{"data": ret})
	})

	server := web.NewService(
		web.Name("xxxyyy"),
		web.Address("0.0.0.0:9999"),
		web.Handler(r),
		web.Registry(reg),
	)

	server.Run()
}
