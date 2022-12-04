package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/micro/v3/service"
	testpb "go-micro-demo/services/test/proto"
	"net/http"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "index",
	})
}

func ServiceOne(c *gin.Context) {

	service := service.New()

	service.Init()

	// 创建微服务客户端
	client := testpb.NewTestService("test", service.Client())

	// 调用服务
	rsp, err := client.Call(c, &testpb.Request{
		Name: c.Query("key"),
	})

	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": 200, "msg": rsp.Msg})
}
