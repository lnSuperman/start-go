package main

import (
	"fmt"
	"strconv"
)

func main() {
	//基本类型转换
	a := int(3.0)
	fmt.Println(a)
	//go语言不支持变量间的隐式类型转换
	var c float32 = 1.43
	//变量与常量隐式转换也需要x.0
	var b int = 7.0
	fmt.Println(b, c)

	d := 6.1
	var e int = int(d)
	fmt.Println(e)
	//字符串转float会报错
	//f := "a"
	//var g int = int(f)
	//fmt.Println(g)
	// i to a
	f := strconv.Itoa(e)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
	//a to i
	g, err := strconv.Atoi(f)
	if err != nil {
		fmt.Println(g)
		fmt.Printf("%T", g)
	}
	//仅1与true、True为true 其他值均为false
	//string to bool
	h, err := strconv.ParseBool("q")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(h)
	// string to float
	// string to float32 -> float64  用bitSize 32 损失精度 返回1.3240000009536743
	// string to float64 -> float64
	i, err := strconv.ParseFloat("1.324", 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	// string TO int  String必须是整数，否则报错
	j, err := strconv.ParseInt("134", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	println(j)
	//将给定格式转为String
	k := strconv.FormatBool(true)
	fmt.Println(k)
	l := strconv.FormatFloat(2.3334, 'E', -1, 64)
	//bitSize 来源类型 32、64
	//fmt输出格式
	//prec 精度
	fmt.Println(l)
	m := strconv.FormatInt(-64, 16)
	//将十进制i转换为base进制并转换为string
	fmt.Println(m)
	n := strconv.FormatUint(64, 16)
	fmt.Println(n)
}
