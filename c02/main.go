package main

import "fmt"

func main() {
	// 单变量
	//var i int
	//i = 20
	//fmt.Println(i)

	//var i int = 20
	//fmt.Println(i)

	//var j = 20
	//fmt.Println(j)

	//j := 100
	//fmt.Println(j)

	//var i type = value
	//var i = value
	//i := value
	// 多变量
	//var a, b, c int = 1, 2, 3
	//fmt.Println(a, b, c)

	// 集合类型
	//var (
	//	a int
	//	b string
	//	c string
	//)
	//a = 10
	//b = "qa"
	//c = "test"
	//fmt.Println(a, b, c)

	//	常量
	//const PI = 3.1415926
	//r := 2.0
	//c := 2 * PI * r
	//fmt.Println(c)
	//
	//const (
	//	unknown        = 0
	//	female         = 1
	//	a       int    = 10
	//	y       string = "test"
	//)
	//fmt.Println(unknown, female, a, y)

	//	const iota 常量计数器
	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)

	const (
		d = iota
		e
		f
		g, h = iota, iota
		i    = iota
	)
	fmt.Println(d, e, f)
	//1.iota 只能在常量中使用
	//2.每一个const常量组之间互不干扰 iota
	//3.没有表达式的常量定义复用上一行的表达式
	//4.从第一行开始iota逐行加一
	//5.同一行iota值相同，代表这里的行数
}
