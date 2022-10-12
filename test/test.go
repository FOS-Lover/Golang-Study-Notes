package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go showMsg("test") // 启动一个协程来运行
	go showMsg("Golang")

	time.Sleep(time.Millisecond * 2000)
	fmt.Println("main end") // 主函数退出 程序就结束了

}
