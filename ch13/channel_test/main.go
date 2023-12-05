package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Consumer(queue chan int) {

	defer wg.Done()
	//方式一 for range
	//for data := range queue {
	//	fmt.Println(data)
	//}
	//方式二 <-
	//data := <-queue
	//fmt.Println(data)

	//方式三
	for {
		data, ok := <-queue
		if !ok {
			break
		}
		fmt.Println(data)
		time.Sleep(time.Second)
	}

}

func main() {

	/*
		channel 提供 通信机制
	*/
	////1.定义channel
	//var msg chan int
	////2.初始化
	////方式一，无缓冲
	////msg = make(chan int)
	////方式二，有缓冲
	//msg = make(chan int, 1)
	////slice map  channel
	//msg <- 1
	//
	//wg.Add(1)
	//go Consumer(msg)
	//msg <- 2
	//msg <- 3
	//close(msg)
	////msg <- 4
	////关闭的channel 不能再次发送数据  已经关闭的channel还可以取数据
	//wg.Wait()

	//无缓冲
	var msg chan int
	msg = make(chan int)
	wg.Add(1)
	go Consumer(msg)
	msg <- 1
	msg <- 2
	close(msg)

	wg.Wait()

}
