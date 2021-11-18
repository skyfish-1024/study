package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex //声明全局互斥锁，防止资源竞争
var wg1 sync.WaitGroup

//定义一个验证是否为素数的函数
func Primenumber(i int, count *int, numbers *[]int) {
	defer wg1.Done()//每检查一次，计数器-1
	for j := 3; j < i; j = j + 2 {
		if i%j == 0 {
			break
		}
		if j == i-2 {
			lock.Lock() //写数据前上锁
			(*numbers) = append((*numbers), i)
			*count++
			lock.Unlock() //写完数据后解锁
		}
	}

}
func main() {
	start := time.Now() //开始时间
	var (
		numbers []int = []int{2, 3} //用于储存素数
		count   int   = 2           //用于计算素数个数(包含2，3，所以初始值为2)
	)
	//runtime.GOMAXPROCS(12)
	wg1.Add(24999)//3~50000的奇数个数，即需要判断的个数，也是计数器的计数值
	//分别对3~50000的奇数进行检查
	for i := 3; i < 50000; i = i + 2 {
		go Primenumber(i, &count, &numbers)//开协程检查是否为素数
	}
	wg1.Wait()//等待子协程全部结束
	fmt.Println(count)//打印素数个数
	end := time.Now()//结束时间
	fmt.Println(end.Sub(start))//计算程序运行的时间

}

//共5133个素数
