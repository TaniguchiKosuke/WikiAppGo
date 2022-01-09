package contorollers

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"WikiAppGo/app/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func loginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

func login(c *gin.Context) {

}

func signupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

func signup(c *gin.Context) {
	var form models.User
	if err := c.Bind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
		c.Abort()
	} else {
		username := c.PostForm("username")
		email := c.PostForm("email")
		password := c.PostForm("password")
		// 登録ユーザーが重複していた場合にはじく処理
		if err := createUser(username, email, password); err != nil {
			errMsg := "This Email is already used"
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": errMsg})
		}
		c.Redirect(302, "/")
	}
}

func createUser(username string, email string, password string) error {
    passwordEncrypt, err := bcrypt.GenerateFromPassword([]byte(email), 12)
	if err != nil {
		log.Println(err)
		return err
	}
    db := models.DbConnect()

	//User idのためのuuidを生成
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return err
	}
	id := uuid.String()

	//新しくUserの作成
	err = db.Create(&models.User{ID: id, Username: username, Email: email, Password: passwordEncrypt}).Error
    if  err != nil {
        return err
    }
    return nil
}

func StartWebServer() {
	router := gin.Default()
	router.Static("/assets", "app/assets/")
	router.LoadHTMLGlob("app/views/*")
	router.GET("/", home)
	router.GET("/user/login", loginPage)
	router.POST("/user/login", login)
	router.GET("/user/signup", signupPage)
	router.POST("/user/signup", signup)
	router.Run()
}