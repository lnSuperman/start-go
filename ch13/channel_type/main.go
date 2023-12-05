package main

import (
	"fmt"
	"sync"
	"time"
)

func Consumer(queue <-chan int) {
	/*
		双向队列
	*/
	defer wg.Done()
	for {
		data, ok := <-queue
		if !ok {
			break
		}
		fmt.Println(data)
		time.Sleep(time.Second)

	}

}

func stock(queue chan<- int) {
	/*
		双向队列
	*/
	defer wg.Done()
	for {
		queue <- 1
		time.Sleep(time.Second)
	}

}

var wg sync.WaitGroup

func main() {
	var msg chan int //双向
	//var msg chan<- int //  单向
	msg = make(chan int)
	wg.Add(1)
	go Consumer(msg) // 双向可以直接转成单向 ，可以装双向参数传递给接收单向参数的方法
	msg <- 1
	//fmt.Println(<-msg)
	close(msg)
	wg.Wait()

}

//func main() {
//	var msg chan<- int //  单向 只可取
//	//var msg <-chan int //  单向 只可放   一般不会用作初始化 用作参数限制
//	msg = make(chan int, 10)
//	msg <- 1
//	msg <- 2
//	data := <-msg
//	fmt.Println(data)
//}
