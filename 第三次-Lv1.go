package main

import (
	"fmt"
	"sync"
)
func A(chC,chA chan int){
	defer wg.Done()//计数-1
	defer close(chA)//最后关闭通道
	for i:= 0; i <10;i++{
		<-chC//若可读，则可继续，用来控制打印顺序
		fmt.Printf("A")
		chA<-1//写入chA
	}
}

func B(chA,chB chan int) {
	defer wg.Done()//计数-1
	defer close(chB)//最后关闭通道
	for i := 0; i < 10; i++ {
		<-chA//若可读，则可继续，用来控制打印顺序
		fmt.Printf("B")
		chB <- 1//写入chB
	}
}

func C(chB,chC chan int){
	defer wg.Done()//计数-1
	defer close(chC)//最后关闭通道
	for i:= 0; i <10;i++{
		<-chB//若可读，则可继续，用来控制打印顺序
		fmt.Printf("C ")
		chC<- 1//写入chC
	}
}

var wg sync.WaitGroup

func main() {
	chA := make(chan int,1)
	chB := make(chan int,1)
	chC := make(chan int,1)
	chC <- 1//开始的标志，写入chC

	go A(chC,chA)
	go C(chB,chC)
	go B(chA,chB)

	wg.Add(3)//计数为3
	wg.Wait()//等待协程结束

}
