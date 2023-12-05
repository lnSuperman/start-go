package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var p *int = new(int)
	//没有初始值 没有内存
	//go编译器知道先申请一个内存空间 这里内存的值全部申请为0
	*p = 10
	fmt.Println(*p)
	//info := map[string]string{}
	var info map[string]string = make(map[string]string)
	//int使用make就不行
	//make同样可以申请内存 make函数一般用于切片 map
	info["c"] = "qa"
	var info2 map[string]string
	//new返回值为指针  make返回指定类型的实例
	if info2 == nil {
		fmt.Println("info2 nil")
	}
	var sliceE []string
	if sliceE == nil {
		fmt.Println("slice nil")
	}
	//map slice nil
	var err error
	if err == nil {
		fmt.Println("err nil")

	}
	//python的None和go的nil类型不一样，None是全局唯一的
	//go语言中nil是唯一用来表示部分类型0值的唯一标识符，它可以代表许多内存布局的值
	fmt.Println(unsafe.Sizeof(sliceE), unsafe.Sizeof(info2))
}
