package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//blog主页
func blog(c *gin.Context){
	Articles,err:= queryArticle()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	LeavingMassages,err:=queryleavingmassage(user.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	//返回数据
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Leavingmassage":LeavingMassages,
		"username":user.Username,
		"Articles":Articles,
	})
}

//login GET
func login(c *gin.Context){
	c.HTML(http.StatusOK,"login.html",nil)
}

//Login POST
func Login(c *gin.Context) {
	//获取form参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	err := LoginCheck(username, password) //检查用户名和密码是否正确
	if err != nil {
		//c.JSON(200,gin.H{
		//	"jieg ":"ghjk",
		//})
		c.HTML(http.StatusOK, "login.html", gin.H{"outcome": err})
	} else {
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/blog/home")
	}
}

var user *User
//初始化user
func CreateUser(){
	user=new(User)
	user.Username="游客"
}
//登录账号密码验证
func LoginCheck(username string,password string)(err error)  {
	//user=new(User)
	sqlStr:="select id,username,password,age,gender from user where username=?"
	err = db.QueryRow(sqlStr,username).Scan(&user.Id,&user.Username, &user.Password, &user.Age,&user.Gender)
	if err!=nil{
		err=errors.New("该用户不存在！")
		fmt.Println(err)
		return
	}else if user.Password!=password{
		err=errors.New("密码错误！")
		fmt.Println(err)
		return err
	}else {
		//token,_:=SetToken(username)
		//fmt.Println("token:",token)
		fmt.Println("登录成功！")
	}
	return
}





//注册GET
func register(c *gin.Context){
	c.HTML(http.StatusOK,"register.html",nil)
}

//注册POST
func Register(c *gin.Context) {
	Username := c.PostForm("username")
	Password := c.PostForm("password")
	Agestr := c.PostForm("age")
	Age,_:=strconv.Atoi(Agestr)
	Gender := c.PostForm("gender")
	security_question:=c.PostForm("security_question")
	err := RegisterCheck(Username, Password, Age, Gender,security_question)
	if err == nil {
		c.HTML(http.StatusOK, "login.html", nil)
	} else { //若存在，则不可以注册
		c.HTML(http.StatusOK, "register.html", gin.H{
			"outcome": err,
		})
	}
}
//注册查重
func RegisterCheck(Username string, Password  string, Age int, Gender string, security_question string) ( err error) {
	if Username==""{
		fmt.Println("账号不能为空")
		return errors.New("账号不能为空")
	}else if Password==""{
		fmt.Println("密码不能为空")
		return errors.New("密码不能为空")
	}
	sqlStr:="insert into user (username,password,age,gender,security_question)values (?,?,?,?,?)"
	_,err=db.Exec(sqlStr,Username,Password,Age,Gender,security_question)
	if err!=nil{
		fmt.Println(err)
		return
	}
	return
}


//back GET
func back(c *gin.Context){
	c.HTML(http.StatusOK,"back.html",nil)
}
//back POST
func Back(c *gin.Context){
	username:=c.PostForm("username")
	security_question:=c.PostForm("security_question")
	U,err:=BackCheck(username,security_question)
	if err!=nil{
		c.JSON(200,gin.H{"结果：":err})
		return
	}
	c.JSON(200,gin.H{
		"username:":U.Username,
		"password:":U.Password,
	})

}
//密保找回
func BackCheck(username string,security_question string)(u User,err error)  {
	user=new(User)
	sqlStr:="select id,username,password,age,gender,security_question from user where username=?"
	err= db.QueryRow(sqlStr,username).Scan(&user.Id,&user.Username, &user.Password, &user.Age,&user.Gender,&user.security_question)
	if err!=nil{
		fmt.Println(err)
		return *user,err
	}else if user.security_question!=security_question{
		err=errors.New("密保错误")
		fmt.Println(err)
		return *user,err
	}
	fmt.Println(*user)
	return *user,nil
}