package main

import "fmt"

func add2(n int) {
	// 拷贝类型
	n += 2
}

// 使用指针修改非作用域内的变量
func add2ptr(n *int) {
	*n += 2
}

func main() {
	n := 5
	add2(n)
	fmt.Println(n) // 5
	// 取址符
	add2ptr(&n)
	fmt.Println(n) // 7
}
