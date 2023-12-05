package main

import "fmt"

func div(a, b int) (int, error) {
	if b == 0 {
		panic("除数不能为0")
	}
	//panic类异常 放到协程里 即使在父协层里有捕获异常，仍会导致主线程挂掉 导致其他协程全部挂掉
	return a / b, nil

}
func main() {
	a := 10
	b := 0
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("异常被捕获到")
		}
		fmt.Println("test")
	}()
	value, err := div(a, b)
	if err != nil {
		println(value)

	}

}
