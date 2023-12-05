package main

import "fmt"

var a int = 20

func main() {
	//局部变量
	a := 10
	fmt.Println(a)
	//fmt.Println(a)
	//在go中字符和字符串不是一种类型 字符类型是单引号 字符串是双引号
	fmt.Printf("%T\n", "慕") // STRING
	fmt.Printf("%T", '慕')   //int32
}
