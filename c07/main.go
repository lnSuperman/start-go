package main

import "fmt"

func main() {
	//var num int
	//_, _ = fmt.Scanf("%d", &num)
	//if num%2 == 0 {
	//	fmt.Printf("%d 是偶数", num)
	//} else {
	//	fmt.Printf("%d 是基数", num)
	//}
	var score int
	_, _ = fmt.Scanf("%d", &score)
	if score >= 80 {
		fmt.Println("成绩优秀")
	} else if score >= 60 && score < 80 {
		fmt.Println("成绩合格")
	} else {
		fmt.Println("成绩不合格 重修")
	}
	var q int
	_, _ = fmt.Scanf("%d", &q)
	if num := q; num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

}
