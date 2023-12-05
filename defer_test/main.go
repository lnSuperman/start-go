package main

import "fmt"

func f1() int {
	i := 10
	defer func() {
		i++
	}()
	return i
}
func f2() *int {
	i := 10
	j := &i
	defer func() {
		*j++
	}()
	return j
}

// return 现将返回的值传递给一个临时变量 先执行defer 再执行return变量
//如果是值传递则无法修改，如果返回指针类的

func main() {
	//fmt.Println("test1")
	//defer println("defer test1")
	//defer println("defer test2")
	//defer println("defer test3")
	////defer 后面只能是函数调用不能是表达式
	////defer 栈 先进后出顺序执行
	//
	//fmt.Println("content")
	//test := func() {
	//	fmt.Println("test1")
	//}
	////压栈时test还没更新
	//defer test()
	//test = func() {
	//	fmt.Println("test2")
	//}

	//x := 10
	//defer func(x int) {
	//	fmt.Println(x)
	//}(x)
	//x++
	// 值也会压栈

	//x := 10
	//defer func(x *int) {
	//	fmt.Println(*x)
	//}(&x)
	//x++
	//// 参数值是指针 指针对应的值更新执行到defer就为内存地址最新的值

	//x := 10
	//defer func() {
	//	fmt.Println(x)
	//}()
	//x++
	////defer 函数未传参 函数内部取到全局变量，则未最后更新的变量值

	fmt.Println(*f2())
	//注册延迟函数 ，defer函数执行顺序已经确定
	//defer 为了取代 try ...

}
