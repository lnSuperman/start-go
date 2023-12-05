package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

//type Info struct {
//	Name string `json:"name,omitempty"`
//	Age  int    `json:"age,omitempty"`
//	Sex  string `json:"sex,omitempty"`
//}

//goland:noinspection ALL
type Info struct {
	Name string `orm:"name,type=varchar(10),max_length=17,min_length=5"`
	Age  int    `json:"age,omitempty" orm:"age,type=int(10),max=70,min=18"`
	Sex  string `orm:"sex,type=varchar(10)"`
	No   string `json:"-"` // 不想序列化
}

func main() {
	info := Info{
		Name: "qa",
		Sex:  "man",
		No:   "test",
	}
	fmt.Println(info.Name, info.Age, info.Sex)
	re, _ := json.Marshal(info)
	fmt.Println(string(re))

	//通过反射识别tag
	t := reflect.TypeOf(info)
	//fmt.Println(t.Name(), t.Kind(), t.NumField())
	for i := 0; i < t.NumField(); i++ {
		//fmt.Println(t.Field(i))
		field := t.Field(i)
		ormTagValue := field.Tag.Get("orm")

		//fmt.Println(tagValue)

		fmt.Println(strings.Split(ormTagValue, ","))

	}
}
