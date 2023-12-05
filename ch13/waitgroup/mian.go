package main

import (
	"fmt"
	"sync"
)

// 子退出主再退出
// add 和 done 必须相等
var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
