package main

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func HomePage(context *gin.Context) {
	context.String(http.StatusOK, "Welcome")
	/*
		context.JSON(http.StatusOK,jsonData)//返回JSON
		context.XML()//返回XML
		context.YAML()//返回YAML
	*/
}

func Default(context *gin.Context) {
	name := context.Param("name")
	sex := context.Params.ByName("sex")
	log.Printf("name:%v,sex:%v\n", name, sex)
}

func PostInfo(context *gin.Context) {
	context.Request.ParseForm()
	log.Printf("name:%v\n", context.PostForm("name"))
	log.Printf("id:%v\n", context.DefaultPostForm("id", "123456"))
	log.Println("Body:", context.Request.Form)
	log.Println("Query name:", context.Query("name")) //依赖于URL后面的参数
	log.Println("Query info:", context.DefaultQuery("info", "vvvvery "))
}

func main() {
	app := gin.Default() //注册实例
	app.GET("/", HomePage)
	app.GET("/default/:name", Default)
	app.POST("/post", PostInfo)
	app.GET("/auth", Auth)
	app.GET("authtest", AuthTest)
	//app.LoadHTMLFiles("view/index.html") //添加HTML文件
	app.Run(":8080")
}

func GetToken(username, password string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
	})
	res, _ := token.SignedString([]byte("secret"))
	return res
}

func Auth(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")
	context.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"token":  GetToken(username, password),
	})

}
func AuthTest(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")
	token := context.Query("token")
	if token == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"info":   "Access Permision Denied",
		})
		return
	} else if token == GetToken(username, password) {
		context.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"info":   "Access Success",
		})
		return
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"info":   "tocken error:Access Permision Denied",
		})
		return
	}
}
