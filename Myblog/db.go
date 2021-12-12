package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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





