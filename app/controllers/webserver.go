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

	c.HTML(
		http.StatusOK,
		"document_detail.html",
		gin.H{
			"document": document,
			"documents": documents,
		})
}

func deleteDocumentConfirm(c *gin.Context) {
	db := models.DbConnect()
	user := getRequestUser(c)
	documents := commons.GetDocumentsList(user, db)
	var document models.Document
	documentId := c.Param("id")
	db.Where("id = ?", documentId).First(&document)
	c.HTML(
		http.StatusOK,
		"delete_document_confirm.html",
		gin.H{
			"documents": documents,
			"document": document,
		})
}

func deleteDocument(c *gin.Context) {
	db := models.DbConnect()
	documentId := c.Param("id")
	var document models.Document
	db.Where("id = ?", documentId).Delete(&document)
	c.Redirect(302, "/")
}

func StartWebServer() {
	getRouter()
}
