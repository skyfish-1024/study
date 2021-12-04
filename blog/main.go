package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//初始化数据库
	err := NewDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	r.GET("blog/home", blogHome)
	//登录
	r.GET("blog/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	//注册
	r.GET("blog/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})
	//发布文章，有个Jwttoken（）中间件
	r.GET("blog/publish",Jwttoken(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "publish.html", nil)
	})
	//登录
	r.POST("blog/login", func(c *gin.Context) {
		//获取form参数
		username := c.PostForm("username")
		userpassword := c.PostForm("password")
		err := login(username, userpassword) //检查用户名和密码是否正确
		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{"outcome": err})
		} else {
			SetToken(username)
			c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/blog/home")
		}
	})
	//注册
	r.POST("blog/register", func(c *gin.Context) {
		Username := c.PostForm("username")
		Userpassword := c.PostForm("password")
		Age := c.PostForm("age")
		Gender := c.PostForm("gender")
		err := register(Username, Userpassword, Age, Gender)
		if err == nil {
			c.HTML(http.StatusOK, "login.html", nil)
		} else { //若存在，则不可以注册
			c.HTML(http.StatusOK, "register.html", gin.H{
				"outcome": err,
			})
		}
	})
	//发布文章，有个Jwttoken（）中间件
	r.POST("blog/publish", func(c *gin.Context) {
		Title:=c.PostForm("title")
		Content:=c.PostForm("content")
		Author:=c.PostForm("author")
		Post_time:=c.PostForm("post_time")
		err:=insertArticle(Title,Content,Author,Post_time,)
		if err!=nil{
			fmt.Println("发布失败",err)
			return
		}else {
			c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/blog/home")
		}
	})
	//删除文章
	r.DELETE("/blog/:articleTitle", func(c *gin.Context) {
		Title:=c.Param("articleTitle")
		err:=deleteArticle(Title)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{"删除结果":"失败！"})
		}
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/blog/home")
	})

	r.Run(":9090")
}


