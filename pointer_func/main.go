package main

import "fmt"

func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c

}
func main() {
	a := 10
	b := 20

	var ip *int
	ip = &a
	*ip = 30
	fmt.Printf("%p ,%d,%d\n", ip, *ip, a)

	swap(&a, &b)
	fmt.Println(a, b)
	//	float array
	//	指针指向数组
	arr := [3]int{1, 2, 3}
	var ips *[3]int = &arr
	var ptrs [3]*int
	fmt.Println(*ips, ptrs)
	//很多都是函数参数的时候指明的类型
	if ips != nil {
		println("为空")

	}
}
