package contorollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func login(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

func signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

func StartWebServer() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("app/views/*")
	router.GET("/", home)
	router.GET("/user/login", login)
	router.GET("/user/signup", signup)
	router.Run()
}