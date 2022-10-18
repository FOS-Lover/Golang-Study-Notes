package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// 解析嵌套指针类型
func f1() {
	b := []byte(`{"Name":"tom","Age":20,"Email":"tom@tom.com", "Parents":["tom","kite"]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	fmt.Println(f)      // map[Age:20 Email:tom@tom.com Name:tom Parents:[tom kite]]
	fmt.Printf("%T", f) // map[string]interface {}Name tom
	for k, v := range f.(map[string]interface{}) {
		fmt.Println(k, v)
	}
}

// 解析嵌套引用类型
func f2() {
	type Person struct {
		Name   string
		Age    int
		Email  string
		Parent []string
	}
	p := Person{
		Name:   "tom",
		Age:    20,
		Email:  "tom@tom.com",
		Parent: []string{"big tom", "big kite"},
	}
	b, _ := json.Marshal(p)
	fmt.Println(string(b))
}

func f3() {
	// 读取json文件
	f, _ := os.Open("test.json")
	defer f.Close()
	dec := json.NewDecoder(f)

	fmt.Println(dec, "\n")
	fmt.Printf("%T \n", dec)

	for {
		var v map[string]interface{}

		err := dec.Decode(&v)
		if err != nil {
			log.Println("Decoder:", err)
			return
		} else {
			fmt.Println("Decoder:", v)
		}
	}
}

func f4() {
	// 写入json文件
	type Person struct {
		Name   string
		Age    int
		Email  string
		Parent []string
	}
	p := Person{
		Name:   "tom",
		Age:    20,
		Email:  "tom@tom.com",
		Parent: []string{"big tom", "big kite"},
	}
	f, _ := os.OpenFile("test.json", os.O_RDWR, 0777)
	defer f.Close()
	enc := json.NewEncoder(f)
	err := enc.Encode(p)
	fmt.Println(err)
}
func main() {
	//f1()
	//f2()
	//f3()
	f4()
}
