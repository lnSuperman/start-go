package main

import (
	"fmt"
	"strconv"
)

type User struct {
	Name string
	Age  int
	Sex  string
}

func printUserInfo(user User) {
	fmt.Println(user.Name, user.Age, user.Sex)

}
func (u User) printUserInfoOne() {
	fmt.Println(u.Name, u.Age, u.Sex)

}
func (u *User) setUserAge(age int) {
	u.Age = u.Age + age
}

type myInt int

func (i myInt) toString() (rInt string) {
	rInt = strconv.Itoa(int(i))
	return

}
func main() {
	u1 := User{
		Name: "liu",
		Age:  18,
		Sex:  "男",
	}
	printUserInfo(u1)
	//函数参数定义
	u1.printUserInfoOne()
	//定义结构体方法
	//User.printUserInfoOne(u1)
	u1.setUserAge(1)
	fmt.Println(u1.Age)
	//User.setUserAge(u1)
	fmt.Println(u1.Age)
	//结构体的接收者有两种形式 1、值传递 2、指针传递
	//结构体的方法只能和结构体在同一个包中
	//int类型不能加方法

	myi := myInt(1)
	fmt.Println(myi.toString())

}
