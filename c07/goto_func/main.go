//package main
//
//import "fmt"
//
//func main() {
//		/* 定义局部变量 */
//		var a int = 10
//
//		/* 循环 */
//	LOOP:
//		for a < 20 {
//			if a == 15 {
//				/* 跳过迭代 */
//				a = a + 1
//				goto LOOP
//			}
//			fmt.Printf("a的值为 : %d\n", a)
//			a++
//		}
//}

//package main
//
//import "fmt"
//
//func main() {
//	var breakAgain bool
//	// 外循环
//	for x := 0; x < 10; x++ {
//		// 内循环
//		for y := 0; y < 10; y++ {
//			// 满足某个条件时, 退出循环
//			if y == 2 {
//				// 设置退出标记
//				breakAgain = true
//				// 退出本次循环
//				break
//			}
//		}
//		// 根据标记, 还需要退出一次循环
//		if breakAgain {
//			break
//		}
//	}
//	fmt.Println("done")
//}

package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 跳转到标签
				goto breakHere
			}
		}
	}
	// 手动返回, 避免执行进入标签
	return
	// 标签
breakHere:
	fmt.Println("done")
}

//err := firstCheckError()
//if err != nil {
//goto onExit
//}
//err = secondCheckError()
//if err != nil {
//goto onExit
//}
//fmt.Println("done")
//return
//onExit:
//fmt.Println(err)
//exitProcess()
