package myweb

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)
func II2() {
	r:=gin.Default()
	r.Handle("GET","/user", func(c *gin.Context) {
		c.String(200,"user api")
	})
	r.Handle("GET","/news", func(c *gin.Context) {
		c.String(200,"news api")
	})

	server:=web.NewService(
		web.Address("0.0.0.0:9999"),
		web.Handler(r),
		)

	server.Run()
}
