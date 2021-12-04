package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)
var db *sql.DB
func NewDB() (err error) {
	//连接数据库
	//数据库信息
	addr := "root:13896764180zy@tcp(127.0.0.1:3306)/Myblog"
	//连接数据库
	db, err = sql.Open("mysql", addr) //检验上面格式是否输入正确
	if err != nil {
		return
	}
	err = db.Ping() //检验是否连接上数据库
	if err != nil {
		return
	}
	//最大连接
	db.SetMaxOpenConns(100)
	//最大空闲
	db.SetMaxIdleConns(100)
	return
}

func blogHome(c *gin.Context) {
	articles, err := queryArticle()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	//返回数据
	c.HTML(http.StatusOK, "home.html", gin.H{
		"username":articles[0].Author,
		"date":    articles,
		"Title":   articles[0].Title,
		"Content": articles[0].Content,
	})
}

var article Article
//获取数据库中所有article信息并储存到article切片中
func queryArticle()(articles []Article,err error ){
	sqlStr:="select id,Title,Content,PointNum,Author,Post_time from article where id >?"
	//err = db.QueryRow(sqlStr).Scan(&article.id,&article.Title, &article.Content, &article.PointNum,&article.Author,&article.Post_time)
	//if err!=nil{
	//	fmt.Println("查询失败！")
	//	return
	//}
	//articles=append(articles,article)
	rows,err:=db.Query(sqlStr,0)
	if err!=nil{
		err=errors.New("查询失败！")
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		er:=rows.Scan(&article.id,&article.Title, &article.Content, &article.PointNum,&article.Author,&article.Post_time)
		if er!=nil{
			return
		}
		articles=append(articles,article)
	}
	return
}

//发布文章
//给数据库添加article信息，即发布
func insertArticle(Title string, Content  string, Author string, Post_time string) ( err error) {
	sqlStr:="insert into article(Title,Content,Author,Post_time)values (?,?,?,?)"
	db.Exec(sqlStr,Title,Content,Author,Post_time)
	if err!=nil{
		err=errors.New("添加失败！")
		fmt.Println(err)
		return
	}
	return
}

//修改文章
//修改数据库中article信息
func updateArticle(Title string, Content  string,Author string, Post_time string,id int) ( err error) {
	sqlStr:="update article Title=?,Content=?,PointNum=?,Author=?,Post_time=? where id=?"
	_,err=db.Exec(sqlStr,Title,Content,Author,Post_time,id)
	if err!=nil{
		err=errors.New("修改失败！")
		fmt.Println(err)
		return
	}
	return
}

//文章点赞
func pointArticle(PointNum int,id int) ( err error) {
	sqlStr:="update article PointNum=? where id=?"
	_,err=db.Exec(sqlStr,PointNum,id)
	if err!=nil{
		err=errors.New("点赞失败！")
		fmt.Println(err)
		return
	}
	return
}

//删除文章
//删除article信息
func deleteArticle(Title string) ( err error) {
	sqlStr:="delete from article where id=?"
	_,err=db.Exec(sqlStr,Title)
	if err!=nil{
		err=errors.New("删除失败！")
		fmt.Println(err)
		return
	}
	return
}

var user *User
//登录账号密码验证
func login(username string,userpassword string)(err error)  {
	user=new(User)
	sqlStr:="select id,username,userpassword,age,gender from user where username=?"
	err = db.QueryRow(sqlStr,username).Scan(&user.id,&user.Username, &user.Userpassword, &user.Age,&user.Gender)
	if err!=nil{
		err=errors.New("该用户不存在！")
		fmt.Println(err)
		return
	}else if user.Userpassword!=userpassword{
		err=errors.New("密码错误！")
		fmt.Println(err)
		return err
	}else {
		token,_:=SetToken(username)
		fmt.Println("token:",token)
		fmt.Println("登录成功！")
	}
	return
}

//注册
func register(Username string, Userpassword  string, Age string, Gender string) ( err error) {
	sqlStr:="insert into user (username,userpassword,age,gender)values (?,?,?,?)"
	db.Exec(sqlStr,Username,Userpassword,Age,Gender)
	if err!=nil{
		err=errors.New("注册失败！")
		fmt.Println(err)
		return
	}
	return
}
