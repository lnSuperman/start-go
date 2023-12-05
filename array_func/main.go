package main

import (
	"fmt"
)

func printArray(a [3]string) {
	a[0] = "test"
	return

}

func main() {
	/*
		大小确定
		类型一致
	*/
	var a [10]string
	var b = [10]string{"django", "scrapy", "tornado"}
	c := [...]string{"django", "scrapy", "tornado"}
	fmt.Println(a, b, c)
	fmt.Println(c[0])
	c[0] = "python"
	fmt.Println(c[0])
	fmt.Println(a, b, c)

	var d = [4]float32{1.0}
	fmt.Println(d)
	f := [...]int{1, 3, 45, 6, 7, 8, 8}
	fmt.Println(f)

	e := [5]int{4: 100}
	//索引4值设置为100 其他为默认值
	fmt.Println(e)
	h := [...]int{1: 1, 4: 7, 10: 4}
	fmt.Println(h)
	//fmt.Println(len(h))
	//	求长度
	//for i, value := range c {
	//	fmt.Println(i,  value)
	//}
	//sum := 0
	//for _, value := range f {
	//	sum += value
	//}
	//fmt.Println(sum)

	var sum int
	for i := 0; i < len(f); i++ {
		sum += f[i]
	}
	println(sum)

	fmt.Printf("%T\n", e)
	fmt.Printf("%T", h)
	//长度不同的数组类型不一样
	//函数参数是数组，数组作为参数，实际调用的时候是值传递
	printArray(c)
	fmt.Println()

}
