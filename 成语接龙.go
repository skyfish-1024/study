package main

import "fmt"
 //截取每个单词的首字母
func FirstNum (a string)string{

	First:=string(a[0])
	return First
}
//截取每个单词的最后一个字母
func LastNum (a string)string{
	Last:=string(a[len(a)-1])
	return Last
}
func main() {
	m:=make([]string,1)//初始化一个m切片
	var w string
	n:="nil"
	fmt.Println("请输入单个依次输入一串单词，以nil结尾：")
	for i:=0;;i++{
		fmt.Scan(&w)//依次输入单词
		if w==n{   //输入结束的条件
			break
		}else {
			m=append(m,w)//把输入的单词存入切片中，从m[1]开始
		}
	}
	fmt.Printf("请输入一个单词的首字母：")
	var key string
		fmt.Scan(&key)
	//检索第一个符合要求的单词并打印
	for i:=1;i<len(m);i++{
		if key==FirstNum(m[i]){
			m[i],m[1]=m[1],m[i]
			fmt.Printf("%v ",m[1])
			break
		}
	}
	//依次检索能接上前一个单词的单词并打印，若接不上则不打印
	for j:=1;j<len(m)-1;j++ {
		for i := j+1; i < len(m); i++ {
			if LastNum(m[j])== FirstNum(m[i]) {
				m[j+1],m[i]=m[i],m[j+1]
				fmt.Printf("%v ",m[j+1])
				break
			}
		}
		if LastNum(m[j])!= FirstNum(m[j+1]){
			break
		}
	}
}