package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string = "test:慕课网"
	//fmt.Println(len(a))
	////返回字节长度 涉及中文
	//runeA := []rune(a)
	//fmt.Println(runeA)
	////` 原样输出内部字符
	var b string = " Qa learn course  Qa in \"慕课网\" "
	if strings.Contains(b, "慕课网") {
		fmt.Println("慕课网")
	}
	fmt.Println(strings.Index(b, "慕课网"))
	//统计出现位置
	fmt.Println(strings.Count(b, "e"))
	//统计出现几次
	fmt.Println(strings.HasPrefix(b, "o"))
	//判断开头
	fmt.Println(strings.HasSuffix(b, "\""))
	fmt.Println(strings.ToUpper(b))
	fmt.Println(strings.ToLower(b))
	//大小写切换
	fmt.Println(strings.Compare(a, b))
	//比较字符串是否相等  第一个字符大于第二个字符 1 反之-1 相等返回0
	fmt.Println(strings.TrimSpace(b))
	//去除前后空格
	fmt.Println(strings.TrimLeft(b, " Q"))
	//去除指定字符
	fmt.Println(strings.Trim("qatestq", "q"))
	//去除左右两侧
	arrayS := strings.Split("test dev rd pm", " ")
	fmt.Println(arrayS)
	//分割
	fmt.Println(strings.Join(arrayS, ","))
	//join
	fmt.Println(strings.Replace(b, "Qa", "Rd", 1))
	//替换

}
