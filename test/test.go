package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int)
var chanStr = make(chan string)

func main() {
	go func() {
		chanInt <- 100
		chanStr <- "Hello"
		defer close(chanInt)
		defer close(chanStr)
	}()
	for {
		select {
		case r := <-chanInt:
			fmt.Println("chanInt:", r)
		case r := <-chanStr:
			fmt.Println("chanStr:", r)
		default:
			fmt.Println("default...")
		}
		time.Sleep(time.Second)
	}
}
