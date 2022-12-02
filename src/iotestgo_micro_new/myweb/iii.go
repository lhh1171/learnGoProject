package myweb

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	//"github.com/micro/go-plugins/registry/consul"
)

func III3() {
	reg := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	r := gin.Default()
	r.Handle("GET", "/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"xxx": "yyy",
			"zzz": "kkk",
		})
	})
	type Obj struct {
		Id   string `form:"id"`
		Name string `form:"name"`
	}
	r.Handle("POST", "/user", func(c *gin.Context) {
		var o Obj
		c.Bind(&o)
		c.JSON(200, gin.H{
			o.Id:  o.Name,
			"xxx": "yyy",
			"zzz": "kkk",
		})
	})

	server := web.NewService(
		web.Name("xxxyyy"),
		web.Address("0.0.0.0:9999"),
		web.Handler(r),
		web.Registry(reg),
	)

	server.Run()
}
