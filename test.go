package main

import "fmt"

type Box struct {
	title   string
	url     string
	setting Setting
}

type Setting struct {
	account  string
	password string
	token    string
	info     Info
}

type Info struct {
	avatar string
	name   string
	age    int
}

func (receiver *Setting) Login(account string, password string) {
	if account == receiver.account && password == receiver.password {
		setToken(receiver)
	}
}

func setToken(setting *Setting) {
	setting.token = "token-true"
}

func main() {
	box := Box{
		title: "test",
		url:   "127.0.0.1",
		setting: Setting{
			account:  "admin",
			password: "admin",
		},
	}

	box.setting.Login("admin", "admin")
	if box.setting.token == "token-true" {
		box.setting.info.avatar = "https://"
		box.setting.info.age = 20
		box.setting.info.name = "tom"
	} else {
		fmt.Println("login error")
	}

	fmt.Println(box)
	fmt.Println(box.setting)
	fmt.Println(box.setting.info)
}
