package contorollers

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"WikiAppGo/app/models"
	"WikiAppGo/app/commons"
)

func getRequestUser(c *gin.Context) models.User {
	session := sessions.Default(c)
	requestUserId := session.Get("UserId")

	db := models.DbConnect()
	var user *models.User
	db.First(&user, "id = ?", requestUserId)
	return *user
}

func index(c *gin.Context) {
	db := models.DbConnect()
	user := getRequestUser(c)
	documents := commons.GetDocumentsList(user, db)
	c.HTML(http.StatusOK, "index.html", gin.H{"documents": documents})
}

func createDocumentPage(c *gin.Context) {
	db := models.DbConnect()
	user := getRequestUser(c)
	documents := commons.GetDocumentsList(user, db)
	c.HTML(http.StatusOK, "create_document.html", gin.H{"documents": documents})
}

func createDocument(c *gin.Context) {
	requestUser := getRequestUser(c)
	db := models.DbConnect()

	//DocumentのprimaryKeyのためのuuidを生成
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
	}
	id := uuid.String()
	title := c.PostForm("title")
	content := c.PostForm("content")

	document := models.Document{
		ID:       id,
		Title:    title,
		Content:  content,
		AuthorID: requestUser.ID,
		Author:   requestUser,
	}
	db.Create(&document)

	c.Redirect(302, "/")
}

func getDocumentDetail(c *gin.Context) {
	documentId := c.Param("id")
	db := models.DbConnect()
	var document models.Document
	user := getRequestUser(c)
	documents := commons.GetDocumentsList(user, db)
	db.Where("id = ?", documentId).First(&document)
	log.Println("############")
	log.Println(document.Title, document.Content)

	c.HTML(
		http.StatusOK,
		"document_detail.html",
		gin.H{
			"document": document,
			"documents": documents,
		})
}

func StartWebServer() {
	getRouter()
}
