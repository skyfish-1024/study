package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type Passage struct {
	Article string
	Author string
	PointNum int
}
var cookiename *string
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
		c.HTML(http.StatusOK,"home.html",gin.H{"username":"游客！",
			"passage":"亲爱的游客你好！美好的一天，从撸代码开始🤣！"})
		c.Abort()//若验证不通过，不再调用后续函数处理
		return
	}
}
func main() {
	cookiename=new(string)
	//初始化一个map，储存用户名和密码
	Records:=make(map[interface{}]interface{})
	Records["test"]="123456"
	var test Passage
	Passages:=make([]interface{},0)
	r:=gin.Default()

	//静态文件预处理
	r.LoadHTMLGlob("templates/*")

	//访问login的GET请求
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",nil)
	})

	//访问register的GET请求
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK,"register.html",nil)
	})

	//访问home的GET请求
	r.GET("/home", AuthMiddleWare(),func(c *gin.Context) {
		c.HTML(http.StatusOK,"home.html",gin.H{
			"username":*cookiename,
			"all":Passages,
		})
	})

	//
	r.GET("/publish",AuthMiddleWare(), func(c *gin.Context) {
		c.HTML(http.StatusOK,"publish.html",nil)
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
				c.SetCookie(*cookiename,"yes",120,"/","localhost",false,true)
				c.HTML(http.StatusOK,"home.html",gin.H{
					"username":username,
					"passage":Passages,
				})
			}else {
				c.HTML(http.StatusOK,"login.html",gin.H{"outcome":"登录失败，密码错误！"})
			}
		}else{
			c.HTML(http.StatusOK,"login.html",gin.H{"outcome":"用户名不存在！"})
		}
	})
	//访问register的POST请求：注册
	r.POST("/register", func(c *gin.Context) {
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		_,ok:=Records[username]//判断用户名是否已经存在
		if !ok{//若不存在，则可以注册
			Records[username]=password
			c.HTML(http.StatusOK,"login.html",nil)
		}else{//若存在，则不可以注册
			c.HTML(http.StatusOK,"register.html",gin.H{
				"outcome":"用户名已存在",
			})
		}
	})
	r.POST("/publish", func(c *gin.Context) {
		passage:=c.PostForm("passage")
		author:=c.PostForm("author")
		test.Article=passage
		test.Author=author
		Passages=append(Passages,test)
		c.HTML(http.StatusOK,"home.html",gin.H{
			"username":*cookiename,
			"passage":test.Article,
			"PointNum":test.PointNum,
			"Author":test.Author,
		})
	})

	r.Run(":9099")
}
