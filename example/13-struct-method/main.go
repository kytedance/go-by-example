package main

import "fmt"

type user struct {
	name     string
	password string
}

// 结构体方法 类似于成员函数
func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

func main() {
	a := user{name: "wang", password: "1024"}
	fmt.Println(a)
	a.resetPassword("2048")
	fmt.Println(a)
	fmt.Println(a.checkPassword("2048")) // true
}
