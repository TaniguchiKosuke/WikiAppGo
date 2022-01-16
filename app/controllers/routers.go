package contorollers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func getRouter() {
	router := gin.Default()
	
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	
	router.Static("/assets", "app/assets/")
	router.LoadHTMLGlob("app/views/**/*")
	
	loginRequired := router.Group("/")
	loginRequired.Use(sessionCheck())
	{
		loginRequired.GET("/", index)
		loginRequired.POST("/logout", logout)
		loginRequired.GET("/create/document/", createDocumentPage)
		loginRequired.POST("/create/document/", createDocument)
		loginRequired.GET("/document/detail/:id/", getDocumentDetail)
		loginRequired.GET("/delete/document/confirm/:id/", deleteDocumentConfirm)
		loginRequired.GET("/delete/document/:id/", deleteDocument) //cannot work with DELETE
		loginRequired.GET("/update/document/form/:id/", updateDocumentPage)
		loginRequired.PATCH("/update/document/:id/", updateDocument)
	}
	
	router.GET("/login", loginPage)
	router.POST("/login", login)
	router.GET("/signup", signupPage)
	router.POST("/signup", signup)
	router.Run(":8000")
}
