package contorollers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"WikiAppGo/app/models"
)

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func loginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func login(c *gin.Context) {
	requestUser := getUser(c.PostForm("email"))

	// DBからハッシュ化されたpasswordを取得
	dbPassword := requestUser.Password
	formPassword := c.PostForm("password")

	//DBから取得したpasswordとformに入力されたpasswordの比較
	if err := bcrypt.CompareHashAndPassword(dbPassword, []byte(formPassword)); err != nil {
		log.Println("Login failed:", err)
		errMsg := "パスワードが一致しません"
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"err": errMsg})
		c.Abort()
	} else {
		log.Println("Login seccussed")
		c.Redirect(302, "/")
	}
}

func getUser(email string) models.User {
	db := models.DbConnect()
	var user models.User

	db.First(&user, "email = ?", email)
	return user
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
			log.Println("email is already used:", err)
			errMsg := "This Email is already used"
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": errMsg})
		}
		c.Redirect(302, "/")
	}
}

func createUser(username string, email string, password string) error {
	passwordEncrypt, err := bcrypt.GenerateFromPassword([]byte(password), 12)
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
	if err != nil {
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
