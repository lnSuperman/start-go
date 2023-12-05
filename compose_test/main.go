package main

import "fmt"

type Detail struct {
	School    string
	ClassName string
	Teacher   string
}
type User struct {
	Name string
	Age  int
	Detail
}

func (d Detail) detailInfo() {
	fmt.Println(d.Teacher, d.ClassName, d.School)

}
func (u User) UserInfo() {
	fmt.Println(u.Name, u.Age, u.School,
		u.Detail.ClassName, u.Teacher)
}

// 匿名嵌套

func main() {
	//go组合实现继承效果
	d1 := Detail{
		School:    "学校",
		ClassName: "班级",
		Teacher:   "老师",
	}
	u1 := User{
		Name:   "学生一",
		Age:    18,
		Detail: d1,
	}
	d1.detailInfo()
	u1.UserInfo()

}
