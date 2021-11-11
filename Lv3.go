package main

import "fmt"

//视频详情结构体
type video struct {
	author        string //作者
	pointNumber   int    //累计获赞数
	coinNumber    int    //累计投币数
	collectNumber int    //收藏次数
	duration      string //视频时长
	postTime      string //发布时间
	introductory  string //简介
}

//发布视频
func post(autherName string) video {
	var first video
	first.author = autherName
	fmt.Println("请输入视频时长：")
	fmt.Scan(&first.duration)
	fmt.Println("请输入发布时间：")
	fmt.Scan(&first.postTime)
	fmt.Println("请输入你的视频简介：")
	fmt.Scan(&first.introductory)
	return first
}

//功能接口
type function interface {
	point()
	collect()
	coin()
	threePoint()
}

//点赞
func (Author *video) point() {
	Author.pointNumber++
}

//收藏
func (Author *video) collect() {
	Author.collectNumber++
}

//投币
func (Author *video) coin() {
	Author.coinNumber++
}

//一键三连
func (Author *video) threePoint() {
	for i := 0; i < 3; i++ {
		Author.pointNumber++
	}
}
func main() {
	var A [100]video //声明了video型数组A来储存创建的结构体
	for i := 0; i < 100; {
		var op string
		fmt.Println("请问需要发布视频吗？（Yes or No）")
		fmt.Scan(&op)
		if op == "Yes" {
			var authorname string
			fmt.Println("请输入你的作者名：")
			fmt.Scan(&authorname)
			A[i] = post(authorname) //创建结构体
			fmt.Printf("作者：%v  赞：%d  投币：%d  收藏：%d次  时长：%vh  发布时间：%v\n简介：%v\n",
				A[i].author, A[i].pointNumber, A[i].coinNumber, A[i].collectNumber,
				A[i].duration, A[i].postTime, A[i].introductory)
			fmt.Printf("请选择以下功能：\n1，点赞\n2.收藏\n3.投币\n4.一健三连\n5.退出\n")
			i++ //用来记录创建的次数
			//通过接口进行操作
			for {
				var n int
				fmt.Scan(&n)
				if n == 5 {
					break
				} else {
					switch n {
					case 1:
						A[i].point()
					case 2:
						A[i].collect()
					case 3:
						A[i].coin()
					case 4:
						A[i].threePoint()
					default:
						break
					}
				}
				fmt.Printf("作者：%v  赞：%d  投币：%d  收藏：%d次 \n",
					A[i].author, A[i].pointNumber, A[i].coinNumber, A[i].collectNumber)
			}
			fmt.Printf("作者：%v  赞：%d  投币：%d  收藏：%d次 \n",
				A[i].author, A[i].pointNumber, A[i].coinNumber, A[i].collectNumber)

		} else {
			break
		}
	}
}
