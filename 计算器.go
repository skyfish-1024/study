package main

import "fmt"

func main() {

	var a,b float64  //定义两个待运算的数
	var c string    //定义运算符
	for {  //for循环，可多次计算
		fmt.Printf("请输入简单计算式(只保留两位有效数字）：\n")//简单提醒输入格式及输出结果格式
		fmt.Scan(&a,&c,&b)
		switch (c) {
		case "+":
			fmt.Printf("%.2f+%.2f=%.2f\n", a, b, a+b)//加法运算
		case "-":
			fmt.Printf("%.2f-%.2f=%.2f\n", a, b, a-b)//减法
		case "*":
			fmt.Printf("%.2f*%.2f=%.2f\n", a, b, a*b)//乘法
		case "/"://除法
			if b == 0 {         //判断分母为零，算式错误
				fmt.Println("分母不能为零！")
				break
			}                   //分母不为零，继续运算
			fmt.Printf("%.2f/%.2f=%.2f\n", a, b, a/b)
		default:
			fmt.Println("输入错误！")
		}
	}
}
