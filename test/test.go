package main

import (
	"fmt"
	"time"
)

func main() {
	// 定时器 4种用法

	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println(time.Now())
	t1 := <-timer1.C // 阻塞的，直到时间到了
	fmt.Println(t1)

	fmt.Println(time.Now())
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Println(time.Now())

	fmt.Println(time.Now())
	time.Sleep(time.Second * 2)
	fmt.Println(time.Now())

	fmt.Println(time.Now())
	<-time.After(time.Second * 2)
	fmt.Println(time.Now())

	// 停止定时器

	timerStop := time.NewTimer(time.Second)
	go func() {
		<-timerStop.C
		fmt.Println("func...")
	}()
	s := timerStop.Stop()
	if s {
		fmt.Println("stop...")
	}
	time.Sleep(time.Second)
	fmt.Println("main end...")

	// 重设定时器
	fmt.Println(time.Now())
	timerReset := time.NewTimer(time.Second * 5)
	timerReset.Reset(time.Second * 2)
	<-timerReset.C
	fmt.Println(time.Now())
}
