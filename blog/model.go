package main

type Article struct {
	id int
	Title    string
	Content  string
	PointNum int
	Author string
	Post_time string
}

type User struct {
	id int
	Username string `form:"username"`
	Userpassword string `form:"password"`
	Age string `form:"age"`
	Gender string `form:"gender"`
}

