package contorollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartWebServer() {
	router := gin.Default()
	router.LoadHTMLGlob("app/views/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"hello": "hello",
		})
	})
	router.Run()
}