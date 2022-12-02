package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func II2() {
	r := gin.Default()
	r.LoadHTMLGlob("/MyProject/goProject/src/iotestgoweb_gin/web/ii.tmpl")
	//r.LoadHTMLFiles("/MyProject/goProject/src/iotestgoweb_gin/web/ii.tmpl", "/MyProject/goProject/src/iotestgoweb_gin/web/ii2.tmpl")
	r.GET("/xxx/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ii/index.html", gin.H{
			"name": "aaabbb",
		})
	})

	err := r.Run("0.0.0.0:9999")
	if err != nil {
		return
	}
}
