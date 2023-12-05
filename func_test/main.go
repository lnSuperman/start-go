package main

import (
	"errors"
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func add2(a, b int) (sum int) {
	// 明确返回值 需要放在括号内指定
	sum = a + b
	return

}
func div(a, b int) (int, error) {
	var err error
	var result int
	if b == 0 {
		err = errors.New("被除数不能为0")
	} else {
		result = a / b
	}
	return result, err

}
func div2(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("被除数不能为0")
	} else {
		result = a / b
	}
	return

}

func main() {
	result, err := div2(4, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
