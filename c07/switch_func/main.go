package main

import "fmt"

func main() {
	/*
			90-100 A
			80-90 B
			70-80 C
		    60-70 D
		    <60 E
	*/
	//var score int
	//_, _ = fmt.Scanf("%d", &score)
	//switch {
	//case score >= 90 && score <= 100:
	//	fmt.Println("a")
	//case score >= 80 && score < 90:
	//	fmt.Println("b")
	//case score >= 70 && score < 80:
	//	fmt.Println("c")
	//case score >= 60 && score < 70:
	//	fmt.Println("d")
	//default:
	//	fmt.Println("e")
	//}
	name := "qa"
	switch name {
	case "qa", "rd":
		fmt.Println("研发部")
	case "pm":
		fmt.Println("产品部")

	}

}
