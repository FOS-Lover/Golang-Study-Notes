package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("%d", rand.Int)
	for i := 0; i < 10; i++ {
		a := rand.Int() // 随机数
		fmt.Printf("%d\n", a)
	}
	fmt.Println("------------")
	for i := 0; i < 10; i++ {
		a := rand.Intn(100) // 一百以内的随机数
		fmt.Printf("%d\n", a)
	}
	fmt.Println("------------")
	for i := 0; i < 10; i++ {
		a := rand.Float32() // 小数的随机数
		fmt.Println(a)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano()) // 以时间作为初始化种子
}
