package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-micro-demo/web/handler"
	"net/http"
)

const addr = ":9000"

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Gin访问成功",
	})
}

func main() {
	r := gin.Default()
	r.Handle("GET", "/", Index)
	r.Handle("GET", "/test-req", handler.ServiceOne)
	if err := r.Run(addr); err != nil {
		fmt.Println("err")
	}
}
