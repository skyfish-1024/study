package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//创建用户
	CreateUser()
	//初始化数据库
	err:=NewDB()
	if err!=nil{
		panic(err)
	}
	r:=gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.GET("blog/home",blog)
	r.GET("blog/login",login)
	r.GET("blog/register",register)
	r.GET("blog/back",back)
	r.GET("blog/publish",publish)
	r.GET("blog/leavemassage",leaveMassage)

	r.POST("blog/login",Login)
	r.POST("blog/register",Register)
	r.POST("blog/back",Back)
	r.POST("blog/publish",Publish)
	r.POST("blog/leavemassage",LeaveMassage)
	r.Run(":9090")
	//insertArticle("熊出没","熊大，光头强又来砍树啦","熊二","12:00")
}
