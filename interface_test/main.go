package main

import "fmt"

type Programmer interface {
	Coding() string
	Debug() string
}
type Designer interface {
	Design() string
}
type Manger interface {
	Programmer
	Designer
	Mange() string
}
type Pythoner struct {
	UiDesigner
	Name string
	Age  int
	Sex  string
}
type UiDesigner struct {
}

func (u UiDesigner) Design() string {
	return " Design !! "

}
func (p Pythoner) Coding() string {
	//fmt.Println("Coding")
	return p.Name + " coding python"

}

func (p Pythoner) Debug() string {
	//fmt.Println("Debug")
	return p.Name + " debug python"

}

func (p Pythoner) Mange() string {
	//fmt.Println("Coding")
	return p.Name + " Mange python"

}

//func (p Pythoner) Design() string {
//	//fmt.Println("Coding")
//	return p.Name + " Design python"
//
//}

type Goer struct {
	Name string
	Age  int
	Sex  string
}

func (goer Goer) Coding() string {
	//fmt.Println("Coding")
	return goer.Name + " coding go"

}

func (goer Goer) Debug() string {
	//fmt.Println("Debug")
	return goer.Name + " debug go"

}

// Pythoner 必须全部实现Programmer Pythoner才是Programmer类型, 否则就不是

func main() {

	//fmt.Println(p1.Coding())
	//fmt.Println(p1.Debug())
	//var pro Programmer = Pythoner{}
	//pro.Coding()
	//pro.Debug()
	var pros []Programmer
	pros = append(pros, Pythoner{},
		Goer{"wang", 19, "man"})
	for _, pro := range pros {
		fmt.Println(pro.Debug())
		fmt.Println(pro.Coding())
	}

	//	go结构体组合一起实现了所有的接口方法
	// 接口本身也支持组合
	var m1 []Manger
	m1 = append(m1, Pythoner{})
	for _, manger := range m1 {

		fmt.Println(manger.Design())

	}
}
