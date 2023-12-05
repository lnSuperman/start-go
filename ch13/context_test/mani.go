package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 监控全局变量
// var stop = false
// 通过chan
//var stop chan bool = make(chan bool)

/*
方案并不统一

*/
//父context 取消 子的也取消
func cpuInfo(ctx context.Context) {
	defer wg.Done()
	ctx2, cancel2 := context.WithCancel(ctx)
	//context.WithDeadline()
	//context.WithTimeout()
	go memoryInfo(ctx2)
	//for {
	//	select {
	//	case <-stop:
	//		fmt.Println("exit get cpu")
	//		return
	//	default:
	//		time.Sleep(time.Second * 2)
	//		fmt.Println("cpu info get done")
	//	}
	//}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("exit get cpu")
			cancel2()
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("cpu info get done")
		}

	}

}

func memoryInfo(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("exit get memory")
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("memory info get done")
		}

	}

}
func main() {
	wg.Add(2)
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	go cpuInfo(ctx)
	//go memoryInfo(ctx)

	//stop <- true
	//close(stop)
	time.Sleep(time.Second * 6)
	//cancel()
	wg.Wait()
	fmt.Println("信息监控完成 ")
}
