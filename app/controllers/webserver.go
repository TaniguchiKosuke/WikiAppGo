package contorollers

import (
	"net/http"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func dbConnect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("cannot open the database")
	}

	return db
}

func getHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func StartWebServer() {
	router := gin.Default()
	router.LoadHTMLGlob("app/views/*")
	router.GET("/", getHome)
	router.Run()
}