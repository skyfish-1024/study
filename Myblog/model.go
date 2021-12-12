package main
//用户
type User struct {
	Id int
	Username string//用户名
	Password string//密码
	Age int//年龄
	Gender string//性别
	security_question string//密保
}
//文章
type Article struct {
	Id int
	Author string//作者
	Title string//标题
	Content string//内容
	PostTime string//发布时间
	PointNum int//点赞数
}
//评论
type Comments struct {
	Id int
	host string     //被评论者
	observer string//评论人
	comment interface{}//评论内容
	PostTime string//评论时间
	PointNum int//点赞数
}
//留言
type Leavingmassage struct {
	Id int
	Host string//被评论者
	Leaver string//留言着
	Comment string//评论内容
	PostTime string//评论时间
	PointNum int//点赞数
}