package main

import "fmt"

func main() {
	// 两种声明和初始化map
	var mapData = map[string]string{
		"name": "Noi-q",
		"age":  "18",
	}
	fmt.Println(mapData) // map[age:18 name:Noi-q]

	var mapData2 = make(map[string]string)
	mapData2["name"] = "Noi-q"
	mapData2["age"] = "20"
	fmt.Println(mapData2) // map[age:20 name:Noi-q]

	fmt.Printf("%T %T\n", mapData, mapData2)

	// 判断key是否存在
	var mapData3 = map[string]string{
		"name": "Noi-q",
	}
	v, ok := mapData3["name"]
	fmt.Println(mapData3, v, ok) // map[name:Noi-q] Noi-q true
}
