package main

import "fmt"

type myErr struct {
}

func (e myErr) Error() string {
	return "错误"

}

func main() {
	//var err error = myErr{}
	//var err error = errors.New("错误")
	var err error = fmt.Errorf("错误：%s", "文件不存在")
	fmt.Println(err)

}
