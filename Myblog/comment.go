package main

import (
	"errors"
	"fmt"
)

//文章评论
func commenter(title string, comment  string, observer string, PostTime string) ( err error) {
	sqlStr:="insert into comments(host,observer,comment,PostTime)values (?,?,?,?)"
	_,err=db.Exec(sqlStr,title,comment,observer,PostTime)
	if err!=nil{
		err=errors.New("评论失败！")
		fmt.Println(err)
		return
	}
	return
}

//评论查询
var comment Comments
func querycomments(host string)(comments []Comments,err error ){
	sqlStr:="select id,observer,comment,PointNum,PostTime from coments where host=?"
	rows,Err:=db.Query(sqlStr,host)
	if Err!=nil{
		fmt.Println("error:",Err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		er:=rows.Scan(&comment.Id,&comment.observer,&comment.comment,&comment.PointNum,&comment.PostTime)
		if er!=nil{
			return
		}
		comments=append(comments,comment)
	}
	return
}
