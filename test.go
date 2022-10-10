package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (receiver Person) eat() {
	fmt.Println("eat")
}
func (receiver Person) sleep() {
	fmt.Println("sleep")
}
func (receiver Person) work() {
	fmt.Println("work")
}
func main() {
	p := Person{
		name: "tom",
		age:  20,
	}
	p.eat()
	p.sleep()
	p.work()
}
