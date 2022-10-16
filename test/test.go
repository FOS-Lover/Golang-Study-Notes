package main

import (
	"bytes"
	"fmt"
)

func f1() {
	// 强制类型转换
	var i int = 100
	var b byte = 10
	b = byte(i)
	i = int(b)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", b)

	var s string = "hello"
	b1 := []byte{1, 2, 3}
	s = string(b1)
	b1 = []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", b1)
}

func f2() {
	// 是否包含
	s := "hello"
	b := []byte(s)
	b1 := []byte("hello")
	b2 := []byte("HELLO")

	b3 := bytes.Contains(b, b1)
	fmt.Println(b3) // true
	b3 = bytes.Contains(b, b2)
	fmt.Println(b3) // false
}

func f3() {
	// 字节切片统计
	s := []byte("hellooo")
	sep1 := []byte("h")
	sep2 := []byte("l")
	sep3 := []byte("o")
	fmt.Println(bytes.Count(s, sep1)) // 1
	fmt.Println(bytes.Count(s, sep2)) // 2
	fmt.Println(bytes.Count(s, sep3)) // 3
}

func f4() {
	// 字节切片重复
	b := []byte("hi")
	fmt.Println(string(bytes.Repeat(b, 1))) // hi
	fmt.Println(string(bytes.Repeat(b, 3))) // hihihi
	fmt.Println(string(bytes.Repeat(b, 9))) // hihihihihihihihihi
}

func f5() {
	// 字节切片替换
	s := []byte("hello,wrold")                           // 字节
	old := []byte("o")                                   // 要被替换的字节
	news := []byte("ee")                                 // 替换的字节
	fmt.Println(string(bytes.Replace(s, old, news, 0)))  // hello,wrold	#第三个参数表示替换几个/几次
	fmt.Println(string(bytes.Replace(s, old, news, 1)))  // hellee,wrold
	fmt.Println(string(bytes.Replace(s, old, news, 2)))  // hellee,wreeld
	fmt.Println(string(bytes.Replace(s, old, news, -1))) // hellee,wreeld
}

func f6() {
	// 字节切片长度和字符串长度
	s := []byte("你好世界")
	r := bytes.Runes(s)
	fmt.Println("字节长度: ", len(s))  // 12
	fmt.Println("字符串长度: ", len(r)) // 4
}

func f7() {
	// 字节切片连接
	s := [][]byte{[]byte("你好"), []byte("世界")} // 二维字节切片
	sep1 := []byte(",")
	fmt.Println(string(bytes.Join(s, sep1)))
	sep1 = []byte("#")
	fmt.Println(string(bytes.Join(s, sep1)))
}

func main() {
	f1()
	f2()
	f3()
	f4()
	f5()
	f6()
	f7()
}
