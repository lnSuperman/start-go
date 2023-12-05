package main

import (
	"fmt"
)

func func1(a, b int) {

}
func func2(params ...int) (sum int) {
	for _, v := range params {
		sum += v
	}
	return

}
func func3(params []int) (sum int) {
	for _, v := range params {
		sum += v
	}
	return

}

type sub func(a, b int) int

// sub 等同于int map
func subImpl(a, b int) int {
	return a - b

}
func filter(score []int, f func(int) bool) []int {
	scoreResult := make([]int, 0)
	for _, value := range score {
		if f(value) {
			scoreResult = append(scoreResult, value)
		} else {
			fmt.Println("err")
			continue
		}
	}
	return scoreResult

}

func main() {
	//result := func2(1, 3, 4, 56, 7)
	//fmt.Println(result)
	slice1 := []int{1, 3, 4, 34, 5, 6}
	fmt.Println(func3(slice1))
	//引用传递 注意不要修改 非值传递
	/*
		省略号的用途 1.函数参数不定长  2.将slice打散 3.定义不定长数组


	*/
	c := [...]int{1, 2, 4}
	fmt.Println(c)
	//匿名函数
	myFunc := func(a, b int) int {
		return a + b
	}
	myFunc(1, 1)
	result := func(a, b int) int {
		return a + b
	}(1, 1)
	fmt.Printf("%T\n", myFunc)
	fmt.Println(result)
	var subQ sub = subImpl
	Q := subQ(1, 2)
	fmt.Println(Q)
	//将函数作为另一个函数的参数
	score := []int{19, 43, 54, 93, 75, 32, 68, 96}

	ruleConf := func(value int) bool {
		if value >= 80 {
			return true
		} else {
			return false
		}
	}

	fmt.Printf("%v", filter(score, ruleConf))
}
