package main

import (
	"fmt"
	"unsafe"
)
import "awesomeProject1/new_pkg"

func main() {
	//go无封装 、继承、多态
	//变量、结构体首字母小写对另一个包不可见
	//结构体定义首字母需要大写
	liu := new_pkg.User{"liu", 18, "https://www.baidu.com"}
	var yang new_pkg.User = new_pkg.User{
		Name: "yang",
		Age:  19,
		Url:  "https://www.baidu.com",
	}
	fmt.Println(liu.Name, yang.Name)
	//如果一个指向结构体的指针 通过结构体指针获取对象的值
	u1 := &new_pkg.User{Name: "zhang", Age: 11, Url: "http://www.baidu.com"}
	fmt.Println((*u1).Age, (*u1).Url, (*u1).Name)
	fmt.Println(u1.Age, u1.Url, u1.Name)
	//go 语法糖 n1.Name  转换成(*n1).Name
	//零值
	u2 := new_pkg.User{}
	fmt.Println(u2, u2.Age)
	//零值初始化结构体
	var u3 new_pkg.User = new_pkg.User{}
	var u4 new_pkg.User
	var u5 *new_pkg.User = new(new_pkg.User)
	fmt.Println(u3.Age, u4.Age, u5.Age)
	// 指针只申明 不赋值 默认值是nil 无法取值
	//结构体是值类型
	n6 := new_pkg.User{"qa", 18, "http://www .baidu.com"}
	n7 := n6
	fmt.Println(n6.Age, n7.Age)
	n6.Age = 22
	fmt.Println(n6.Age, n7.Age)
	//结构体的大小，占用内存的大小，可以使用sizeof来查看对象占用类型
	fmt.Println(unsafe.Sizeof(1))
	/*
		type string struct {
			Data uintptr    //指针占用8位
		    Len int   // 长度64系统占8位长度
				}
	*/
	fmt.Println(unsafe.Sizeof("liu"))
	fmt.Println(unsafe.Sizeof(n6))
	//	string 16 * 2 +  int 8

	type slice struct {
		array unsafe.Pointer //底层数组的地址
		len   int
		cap   int
	}

	s1 := []string{"name", "age", "sex", "dict"}
	fmt.Println(unsafe.Sizeof(s1))
	//切片占用永远都是24
	m1 := map[string]string{
		"name": "liu",
		"age":  "12",
		"text": "test test",
	}
	fmt.Println(unsafe.Sizeof(m1))
	//map 永远都是8
}
