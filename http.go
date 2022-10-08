package main

import "fmt"

type network struct {
	HTTP   string
	HTTPS  string
	SOCKET string
}

func connectHTTP() {
	networkData := network{"http://baidu.com", "https://baidu.com", "socket://127.0.0.1"}
	fmt.Println(networkData)
}

func main() {
	connectHTTP()
}
