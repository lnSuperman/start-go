package main

import (
	"fmt"
	"sort"
)

type Course struct {
	Name  string
	Price int
	Url   string
}

type Courses []Course

func (c Courses) Len() int {
	return len(c)

}

func (c Courses) Less(i, j int) bool {
	return c[i].Price < c[j].Price

}

func (c Courses) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]

}
func main() {
	Courses := Courses{
		Course{"python", 300, ""},
		Course{"test", 400, ""},
		Course{"java", 200, ""},
		Course{"go", 500, ""},
	}

	//c := Course{Name:}
	sort.Sort(Courses)
	//fmt.Printf("%v", Courses)
	for i, cours := range Courses {
		fmt.Println(i, cours)

	}

}
