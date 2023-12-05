package main

import (
	"fmt"
	"sync"
	"time"
)

/*
锁 资源竞争
*/
var total int
var wg sync.WaitGroup
var lock sync.Mutex

// Mutex 互斥锁 同步问题  性能问题
// 读写锁

var RWLock sync.RWMutex

func Read() {
	defer wg.Done()
	RWLock.RLock()
	fmt.Println("读...")
	time.Sleep(time.Second)
	fmt.Println("读成功")
	RWLock.RUnlock()
}
func White() {
	defer wg.Done()
	RWLock.Lock()
	fmt.Println("修改...")
	time.Sleep(time.Second * 10)
	fmt.Println("修改成功")
	RWLock.Unlock()

}

func Add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		total = total + 1
		lock.Unlock()

	}

}
func Sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		total = total - 1
		lock.Unlock()
	}

}

//func main() {
//	wg.Add(2)
//	go Add()
//	go Sub()
//	wg.Wait()
//	fmt.Println(total)
//
//}

func main() {
	wg.Add(6)
	for i := 0; i < 5; i++ {
		go Read()
	}
	for i := 0; i < 1; i++ {
		go White()
	}

	wg.Wait()
}
