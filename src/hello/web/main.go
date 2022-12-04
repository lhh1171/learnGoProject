package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	if err := r.Run(addr); err != nil {
		fmt.Println("err")
	}
}
