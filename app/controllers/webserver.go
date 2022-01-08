package contorollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func StartWebServer() {
	router := gin.Default()
	router.LoadHTMLGlob("app/views/*")
	router.GET("/", getHome)
	router.Run()
}