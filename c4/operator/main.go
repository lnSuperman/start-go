package main

import (
	"fmt"
	"strconv"
)

func main() {
	//a := 12
	//b := 10
	//fmt.Println(a+b, a-b, a*b, a/b, a%b)
	////除法取整
	//a++
	//fmt.Println(a)
	//a--
	//fmt.Println(a)
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if a || b {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if a && b {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(a && b) {
		fmt.Printf("第四行 - 条件为 true\n")
	}

	//	位运算
	//& 与 | 或 ^ 异或
	// & 两个都真才为真
	// | 有一个为真就为真
	//^ 两个真假不同则为真
	var data int
	var err error
	if data, err = strconv.Atoi("12"); err != nil {
		fmt.Println(err)

	}
	fmt.Println(data)

}
