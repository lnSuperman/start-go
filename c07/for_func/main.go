package main

import "fmt"

func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%d", i)
	//
	//}
	//sum := 0
	//for i := 1; i <= 10; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	//for {
	//	println("死循环")
	//}

	//j := 0
	//for j < 10 {
	//	fmt.Println("test")
	//	j++
	//}
	//name := "liu"
	//for _, value := range name {
	//	fmt.Printf("%c ", value)
	//}
	a := "test:慕课网"
	//fmt.Printf("%c\n", a[0])
	//fmt.Printf("%c\n", a[7])
	//for i := 0; i < len(a); i++ {
	//	fmt.Printf("%c\n", a[i])
	//
	//}
	b := []rune(a)
	for i := 0; i < len(b); i++ {
		//println(a[i])
		fmt.Printf("%c", b[i])
	}

}
