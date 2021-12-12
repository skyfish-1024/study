package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//leavemassage GET
func leaveMassage(c *gin.Context){
	c.HTML(http.StatusOK,"leavemassage.html",nil)
}


//leavemassage POST
func LeaveMassage(c *gin.Context){
	Host := c.PostForm("Host")
	Comment := c.PostForm("Content")
	Leaver := c.PostForm("Leaver")
	PostTime := c.PostForm("PostTime")
	err:=leavemassage(Host,Leaver,Comment,PostTime)
	if err!=nil{
		fmt.Println("err:", err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/blog/home")
}

//用户留言
func leavemassage(host string,leaver string,comment interface{},PostTime string)(err error){
	sqlStr:="insert into leavingmassage (host,leaver,comment,PostTime)values(?,?,?,?)"
	_,err=db.Exec(sqlStr,host,leaver,comment,PostTime)
	if err!=nil{
		fmt.Println(err)
		fmt.Println(errors.New("留言失败！"))
		return
	}
	return
}



//留言查询
var leavingmassage Leavingmassage
func queryleavingmassage(host string)(Leavingmassages []Leavingmassage,err error ){
	sqlStr:="select id,leaver,comment,PointNum,PostTime from leavingmassage where host like ?"
	rows,Err:=db.Query(sqlStr,host)
	if Err!=nil{
		fmt.Println("error:",Err,"bbbbb")
		return
	}
	defer rows.Close()
	for rows.Next(){
		er:=rows.Scan(&leavingmassage.Id,&leavingmassage.Leaver,&leavingmassage.Comment,&leavingmassage.PointNum,&leavingmassage.PostTime)
		if er!=nil{
			return
		}
		Leavingmassages=append(Leavingmassages,leavingmassage)
	}
	return
}
