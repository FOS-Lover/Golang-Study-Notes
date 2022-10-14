package main

import "os"

func write() {
	// 字节写入
	// os.RDWR 读写  os.O_APPEND 追加  os.O_TRUNC 覆盖
	f, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_TRUNC, 777)
	f.Write([]byte("test"))
	f.Close()
}

func writeString() {
	// 字符串写入
	f, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND, 777)
	f.WriteString("Hello World")
	f.Close()
}

func writeAt() {
	// 字节偏数组写入
	f, _ := os.OpenFile("test.txt", os.O_RDWR, 777)
	f.WriteAt([]byte("aaa"), 3)
	f.Close()
}

func main() {
	//write()
	//writeString()
	writeAt()
}
