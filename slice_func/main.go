package main

func replace(a []string) {
	a[0] = "rd"

}
func main() {
	//方式一
	//var names []string
	//names := []string{"liu", "yang", "qa", "jon", "qne", "qda"}
	//fmt.Println(names)
	//fmt.Printf("%T", names)
	//切片为什么需要传长度
	//方式二
	//name := make([]string, 5)
	//fmt.Println(len(name))
	//fmt.Printf("%T", name)
	//	方式三
	//var a = [5]string{"django", "scrapy", "tornado"}
	//subA := a[1:3]
	//fmt.Println(subA)
	//fmt.Printf("%T", subA)
	////方式四
	//b := *new([]int)
	//c := new([]int)
	//fmt.Printf("%T\n", b)
	//fmt.Printf("%T", c)
	//引用传递
	//replace(names)
	//fmt.Println(names)
	//namesC := names[1:4]
	//fmt.Printf("%T, %v\n", namesC, namesC)
	//namesC = append(namesC, "lon", "oon")
	//fmt.Println(namesC)
	////copy 时目标对象长度需要设置好
	//namesD := make([]string, 4)
	//copy(namesD, namesC)
	//fmt.Println(namesC, namesD)
	//appendedName := []string{"lon", "oon"}
	//namesC = append(namesC, appendedName...)
	////函数参数传递规则
	//fmt.Println(namesC)
	//
	//var a = [5]string{"django", "scrapy", "tornado", "python", "java"}
	//aSlice := a[:]
	//aNew := append(aSlice[:1], aSlice[2:]...)
	//fmt.Printf("%v", aNew)
	//现象一
	//a := make([]int, 0)
	//b := []int{1, 2, 4}
	//fmt.Println(copy(a, b))
	//fmt.Println(a)
	////	现象二
	//c := b[:]
	//c[0] = 8
	//fmt.Println(b)
	//fmt.Println(c)
	//// 现象三
	//c = append(c, 9)
	//fmt.Println(b)
	//fmt.Println(c)
	////append 没有影响b 原来的数组
	//c[0] = 9
	//fmt.Println(b)
	//fmt.Println(c)
	////append 函数之后修改不影响 原来的数组
	////现象四
	//fmt.Println(len(c))
	//fmt.Println(cap(c))
	//容量函数
	//长度和容量的意义
	//d := make([]int, 5)
	//fmt.Println(len(d))
	//fmt.Println(cap(d))
	//使用make初始化 len和cap相同

	//e := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	//sE := e[2:4]
	//sEE := e[3:7]
	//for index, value := range sE {
	//	fmt.Printf("index %d, value %d\n", index, value)
	//}
	//e[2] = 7
	//e[3] = 10
	//e[4] = 8
	//sE = append(sE, 1, 5, 4, 7, 9, 0)
	//sEE = append(sEE, 7)
	//fmt.Println(e, sE, sEE)
	//fmt.Printf("len %d , cap %d", len(sE), cap(sE))
	////1,2,4,8,16  2倍<1024
	//// 1.25  >1024
	////基于数组 会影响原来的数组

	//	make 遇到append
	//f := make([]int, 0)
	//var g []int
	//fmt.Printf("len %d , cap %d\n", len(f), cap(f))
	//f = append(f, 1)
	//fmt.Printf("len %d , cap %d\n", len(f), cap(f))
	//fmt.Println(f)
	//fmt.Printf("len %d , cap %d\n", len(g), cap(g))
	//fmt.Println(g)

}
