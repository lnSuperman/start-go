package main

import "fmt"

type User struct {
	Name string
}

func Print(x interface{}) {

	//if v, ok := x.(int); ok {
	//	fmt.Printf("整数 %d", v)
	//}
	//if v, ok := x.(string); ok {
	//	fmt.Printf("字符串 %s", v)
	//
	//}
	switch v := x.(type) {
	case int:
		fmt.Printf("整数 %d\n", v)
	case string:
		fmt.Printf("字符串 %s\n", v)
	}

}

func main() {
	//var i interface{}
	////fmt.Println(i)
	////空接口类似于java和python object
	//i = User{}
	////fmt.Println(i)
	/*
				1、可以把任何类型赋值给空接口变量
			    2、参数传递
		        3、空接口可以定义为map的值
	*/
	//var a int = 1
	//var b string = "test"
	//c := map[string]string{}
	//d := make([]string, 0)
	//Print(a)
	//Print(b)
	//Print(c)
	//Print(d)
	//e := make(map[string]interface{})
	//e["name"] = "test"
	//e["age"] = 18
	//e["weight"] = 75.2

	Print(1)
	Print([]string{})

}
