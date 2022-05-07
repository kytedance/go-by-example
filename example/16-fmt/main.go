package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	s := "hello"
	n := 123
	p := point{1, 2}
	fmt.Println(s, n) // hello 123
	fmt.Println(p)    // {1 2}

	// 可以使用 %v 来打印任意类型的值
	fmt.Printf("s=%v\n", s) // s=hello
	fmt.Printf("n=%v\n", n) // n=123
	fmt.Printf("p=%v\n", p) // p={1 2}
	// 更加详细的打印
	fmt.Printf("p=%+v\n", p) // p={x:1 y:2}
	// 进一步详细，包括类型
	fmt.Printf("p=%#v\n", p) // p=main.point{x:1, y:2}

	f := 3.141592653
	fmt.Println(f) // 3.141592653
	// 保留两位小数
	fmt.Printf("%.2f\n", f) // 3.14
}
