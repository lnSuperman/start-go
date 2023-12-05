package main

import "fmt"

func main() {
	//map --> dict
	m1 := map[string]string{
		"m1": "v1",
	}
	fmt.Printf("%v\n", m1)
	m2 := make(map[string]string)
	m2["m2"] = "v2"
	fmt.Println(m2)
	m3 := map[string]string{}
	fmt.Println(m3)
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	if a == b {
		fmt.Println("eq")
	}

	m0 := map[[3]int]string{}
	fmt.Println(m0)
	mapOne := map[string]string{"key1": "value1"}
	mapOne["key1"] = "keysValue1"
	mapOne["key2"] = "keysValue2"
	fmt.Println(mapOne)
	fmt.Println(mapOne["key1"])
	v, ok := mapOne["key3"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println(ok)
	}
	delete(mapOne, "key3")
	//不报错
	fmt.Println(mapOne)
	for key, value := range mapOne {
		fmt.Println(key, value)
	}
	//list

}
