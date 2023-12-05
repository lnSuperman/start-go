package main

import (
	"fmt"
	"time"
)

func Print() {
	fmt.Println("启动协程打印文本")
}

func main() {
	for i := 0; i < 5; i++ {
		//go Print()
		//i := i
		go func(i int) {
			for {
				fmt.Println(i)
				time.Sleep(time.Second)
			}
		}(i)
	}
	time.Sleep(time.Second * 2)
	//主死从随

}
