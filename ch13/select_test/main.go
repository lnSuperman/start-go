package main

import (
	"fmt"
	"time"
)

//func main() {
//	/*
//				select,作用channel之上 io 多路复用
//			select 会随机公平选择 一个case语句执行
//		select 应用场景一：timeout超时机制
//	*/
//	//var ch1 chan int
//	//var ch2 chan int
//	ch1 := make(chan int, 1)
//	ch2 := make(chan int, 1)
//	ch1 <- 1
//	ch2 <- 2
//	select {
//	case data := <-ch1:
//		fmt.Println(data)
//	case data := <-ch2:
//		fmt.Println(data)
//	}
//}

//func main() {
//	timeout := false
//	go func() {
//		time.Sleep(time.Second * 2)
//		timeout = true
//
//	}()
//	for {
//		if timeout {
//			fmt.Println("结束")
//			break
//		}
//		time.Sleep(time.Microsecond * 10)
//	}
//}

func main() {
	timeout1 := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 5)
		timeout1 <- true
	}()
	timeout2 := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 1)
		timeout2 <- true
	}()
	select {
	case <-timeout1:
		fmt.Println("5结束")
	case <-timeout2:
		fmt.Println("1结束")
		//default:
		//	fmt.Println("继续下一次")
		//检测阻塞 如阻塞 case均未执行 执行default代码
	}
}
