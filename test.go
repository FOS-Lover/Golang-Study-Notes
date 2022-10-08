package main

import "fmt"

// 参数变长
func f1(args ...int) {
	for i, arg := range args {
		fmt.Println(i, arg)
	}
}

// 固定参数和参数变长
func f2(name string, age int, args ...int) {
	fmt.Println(name, age, args)
}

// 验证浅拷贝和深拷贝
// 浅拷贝会生成副本不会影响原来的实参参数值
// 深拷贝会拷贝值的指针，会影响双向数据
// map slice interface channel

// 浅拷贝
func f3(a int) {
	a = 100
	fmt.Println(a) // 100
}

// 深拷贝
func f4(a []int) {
	a[0] = 10000
	fmt.Println(a) // [10000 2 3]
}

func main() {
	f1(1, 2, 3, 4, 5)
	f2("tom", 20, 123, 12313212, 123123123)

	test := 10
	f3(test)
	fmt.Println(test) // 10

	test2 := []int{1, 2, 3}
	f4(test2)
	fmt.Println(test2) // [10000 2 3]
}
