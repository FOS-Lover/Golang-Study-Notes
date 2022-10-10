package main

import "fmt"

type Music interface {
	playMusic()
}

type Video interface {
	playVideo()
}

type Mobile struct{}

func (receiver Mobile) playMusic() {
	fmt.Println("play music")
}

func (receiver Mobile) playVideo() {
	fmt.Println("play video")
}

type Person interface {
	read()
}

type A struct{}
type B struct{}

func (receiver A) read() {
	fmt.Println("A")
}
func (receiver B) read() {
	fmt.Println("B")
}

func main() {
	var m Music = new(Mobile)
	m.playMusic()
	var v Video = new(Mobile)
	v.playVideo()

	var perA Person = new(A)
	var perB Person = new(B)
	perA.read()
	perB.read()
}
