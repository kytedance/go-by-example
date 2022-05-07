package main

import "fmt"

// golang 形参类型后置
func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

// golang 原生支持返回多个值，一般第一个是返回值，第二个是错误信息
func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func main() {
	res := add(1, 2)
	fmt.Println(res) // 3

	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok) // A True
}
