package main

import (
	"fmt"
	"math"
)

func main() {
	name := "bobby"
	age := 18
	//fmt.Println("name:" + name + ",age:" + strconv.Itoa(age))
	//fmt.Printf("name:%v, age:%v\n", name, age)
	////原样输出
	//fmt.Printf("name:%#v, age:%#v\n", name, age)
	////字符串加引号
	//fmt.Printf("name:%T, age:%T\n", name, age)
	////显示类型
	//fmt.Printf("name:%s, age:%d\n", name, age)
	//fmt.Printf("name:%s, age:%+4d\n", name, age)
	////+ 必须显示正负符号  4显示位数 Pad空格(宽度为4，右对齐
	//fmt.Printf("name:%s, age:%b\n", name, age)
	//fmt.Printf("name:%s, age:%o\n", name, age)
	////8进制
	//fmt.Printf("name:%s, age:%x\n", name, age)
	//16进制
	//data := 65
	desc := fmt.Sprintf("name:%6s, age:%d\n", name, age)
	fmt.Println(desc)
	//fmt.Printf("%c", data)
	//unicode 转字符
	//fmt.Printf("%q", data)
	//增加引号

	fmt.Printf("%e\n", math.Pi)
	//科学技术法
	fmt.Printf("%f\n", math.Pi)
	//十进制展示

	//输入输出
	var n string
	var a int
	fmt.Println(" input name and age")
	_, err := fmt.Scanln(&n, &a)
	if err != nil {
		return
	}
	fmt.Println(n, a)

	//var n string
	//var a int
	//fmt.Println("input name and age")
	//_, err := fmt.Scanf("%s %d", &n, &a)
	//if err != nil {
	//	return
	//}
	//fmt.Println(n, a)
}
