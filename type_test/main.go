package main

import "fmt"

func main() {
	//1.定义别名
	type myByte = byte
	var a myByte
	fmt.Printf("%T\n", a)
	//2.基于已有类型定义定义新的类型
	type myInt int
	var b myInt
	fmt.Printf("%T\n", b)

	//3.定义结构体
	type Course struct {
		name  string
		price int
	}
	//4.定义接口
	type Callable interface {
	}

	//5.定义函数别名
	type handle func(str string)
}
