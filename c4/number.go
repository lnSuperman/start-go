package main

import (
	"fmt"
	"math"
)

func main() {
	//布尔类型
	var a bool = true
	var b bool = false
	//有符号8位整数 -128~127
	var c int8 = 1
	var d int16 = 1
	var e int32 = 1
	var f int64 = 1
	//无符号8位证书 0-255
	var g uint8 = 1
	//取决于机器的位数 一般不用
	var i int = 1
	//有符号32位
	var j float32 = 1.1
	fmt.Println(a, b, c, d, e, f, g, i, j)
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxFloat32)
	//为什么64位float与64位int存储差距大，底层存储不一致，64位最大数与精度都比32位高
	h := 71.23
	//默认选择64
	fmt.Printf("%T\n", h)
	k := 23
	fmt.Printf("%T\n", k)
	//byte uint rune
	//	rune = int32
	//	byte uint8
	var l byte = 10
	var q rune = 10
	fmt.Println(l, q)

	o := 'o'
	fmt.Printf("%T\n", o+1)
	fmt.Printf("o+1=%c", o+1)

}
