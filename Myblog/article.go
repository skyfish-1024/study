package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var article Article
//var Articles []Article
//获取数据库中所有article信息并储存到article切片中
func queryArticle() (articles []Article,err error) {
	sqlStr := "select id,author,title,content,PointNum,PostTime from article where id >?"
	rows, Err := db.Query(sqlStr, 0)
	if Err != nil {
		fmt.Println("error:", Err)
		return
	}
	defer rows.Close()
	//var articles []Article
	for rows.Next() {
		er := rows.Scan(&article.Id, &article.Author, &article.Title, &article.Content, &article.PointNum, &article.PostTime)
		if er != nil {
			return
		}
		articles = append(articles, article)
	}
	return
}

//发布文章
//publish GET
func publish(c *gin.Context) {
	c.HTML(http.StatusOK, "publish.html", nil)
}

//publish POST
func Publish(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	auther := c.PostForm("auther")
	posttime := c.PostForm("post_time")
	if title == "" || content == "" {
		c.JSON(200, gin.H{"注意:": "题目或内容不能为空"})
		return
	} else {
		Articles,_:=queryArticle()
		//fmt.Println("Articles:",Articles)
		for _, Art := range Articles {
			if title == Art.Title {
				c.JSON(200,gin.H{"注意:":"该标题已经被使用"})
				return
			}
		}
		err := insertArticle(title, content, auther, posttime)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/blog/home")
	}
}

//给数据库添加article信息
func insertArticle(title string, content string, author string, PostTime string) (err error) {
	sqlStr := "insert into article(title,content,author,PostTime)values (?,?,?,?)"
	_, err = db.Exec(sqlStr, title, content, author, PostTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

//修改文章
//修改数据库中article信息
func updateArticle(Title string, Content string, Author string) (err error) {
	sqlStr := "update article title=?,content=?,author=? where title=?"
	_, err = db.Exec(sqlStr, Title, Content, Author, Title)
	if err != nil {
		err = errors.New("修改失败！")
		fmt.Println(err)
		return
	}
	return
}

//文章点赞
func pointArticle(id int) (err error) {
	sqlStr1 := "select PointNum from article where id=?"
	var PointNum int
	err = db.QueryRow(sqlStr1, id).Scan(PointNum)
	if err != nil {
		err = errors.New("点赞失败！")
		fmt.Println(err)
		return err
	}
	sqlStr := "update article PointNum=? where id=?"
	_, err = db.Exec(sqlStr, PointNum+1, id)
	if err != nil {
		err = errors.New("点赞失败！")
		fmt.Println(err)
		return err
	}
	return
}

//删除文章
//删除article信息
func deleteArticle(Title string) (err error) {
	sqlStr := "delete from article where title=?"
	_, err = db.Exec(sqlStr, Title)
	if err != nil {
		err = errors.New("删除失败！")
		fmt.Println(err)
		return
	}
	return
}
