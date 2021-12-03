package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//定义中间件，用cookie记住登录状态
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取客户端cookie并校验
		cookie,err:=c.Cookie(*cookiename)
		if err==nil{
			if cookie=="yes"{
				c.Next()
				return
			}
		}
		//异常，返回错误
		c.JSON(http.StatusUnauthorized,gin.H{"code":200, "message":"游客你好！"})
		c.Abort()//若验证不通过，不再调用后续函数处理
		return
	}
}
var cookiename *string
func main() {
	cookiename=new(string)
	//初始化一个map，储存用户名和密码
	Records:=make(map[interface{}]interface{})
	Records["test"]="123456"//测试账号和密码
	r:=gin.Default()

	//静态文件预处理
	r.LoadHTMLGlob("templates/*")
	r.Static("/statics","./statics")

	//http://localhost:9097/home测试cookie
	r.GET("/home",AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200,gin.H{"code":200, "message":*cookiename+"你好！"})
	})

	//访问login的GET请求
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",nil)
	})

	//访问register的GET请求
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK,"register.html",nil)
	})
	//访问login的POST请求：登录
	r.POST("/login", func(c *gin.Context) {
		//获取form参数
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		cookiename=&username
		date,ok:=Records[username]//判断用户是否存在
		if ok==true{
			if date==password{//判断密码是否正确
				c.SetCookie(*cookiename,"yes",60,"/","localhost",false,true)
				c.HTML(http.StatusOK,"index.html",gin.H{
					"Name":username,
					"Password":password,
					"all":Records,
					"结果":"登录成功",
				})
			}else {
				c.String(200,"密码错误！请刷新页面。")
			}
		}else{
			c.String(200,"用户不存在！请刷新页面。")
		}
	})
	//访问register的POST请求：注册
	r.POST("/register", func(c *gin.Context) {
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		_,ok:=Records[username]//判断用户名是否已经存在
		if !ok{//若不存在，则可以注册
			Records[username]=password
			c.HTML(http.StatusOK,"register.html",gin.H{
				"outcome1":true,
			})
		}else{//若存在，则不可以注册
			c.HTML(http.StatusOK,"register.html",gin.H{
				"outcome2":true,
			})
		}
	})

	r.Run(":9097")

}