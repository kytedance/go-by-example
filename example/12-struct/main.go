package main

import "fmt"

// 结构体，带类型字段集合
type book struct {
	id     int
	name   string
	author string
	price  float32
}

type user struct {
	name     string
	password string
}

func main() {
	a := user{name: "wang", password: "1024"}
	b := user{"wang", "1024"}
	c := user{name: "wang"}
	c.password = "1024"
	var d user
	d.name = "wang"
	d.password = "1024"

	fmt.Println(a, b, c, d)                 // {wang 1024} {wang 1024} {wang 1024} {wang 1024}
	fmt.Println(checkPassword(a, "haha"))   // false
	fmt.Println(checkPassword2(&a, "haha")) // false

	book1 := book{id: 1, name: "Go语言", author: "张三", price: 100}
	book2 := book{id: 2, name: "Python语言", author: "李四", price: 200}
	book3 := book{id: 3, name: "Java语言", author: "王五", price: 300}
	fmt.Println(book1, book2, book3)
	book1.price = 100.1
	fmt.Println(book1, book2, book3)
	var emptyBook book
	fmt.Println(emptyBook)

}

func checkPassword(u user, password string) bool {
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	// 使用指针可以对于结构体进行修改，也可以减少大结构体拷贝开销
	return u.password == password
}
