package main

import "fmt"

func test() {
	name := "tom"
	age := "20"

	f := func() string {
		return name + age
	}
	msg := f()
	fmt.Println(msg)
}

func main() {
	// 匿名函数
	add := func() string {
		return "Hello World"
	}
	a := add()
	fmt.Println(a)

	// 自调用函数
	b := func(a int, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(b)

	test()
}
