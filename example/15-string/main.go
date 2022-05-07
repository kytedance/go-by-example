package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	// 包含
	fmt.Println(strings.Contains(a, "ll")) // true
	// 对于某个字符串计数
	fmt.Println(strings.Count(a, "l")) // 2
	// 前缀与后缀
	fmt.Println(strings.HasPrefix(a, "he"))  // true
	fmt.Println(strings.HasSuffix(a, "llo")) // true
	// 查找位置
	fmt.Println(strings.Index(a, "ll")) // 2
	// 连接字符串
	fmt.Println(strings.Join([]string{"he", "llo"}, "-")) // he-llo
	// 重复
	fmt.Println(strings.Repeat(a, 2)) // hellohello
	// 替换
	fmt.Println(strings.Replace(a, "e", "E", -1)) // hEllo
	// 分割
	fmt.Println(strings.Split("a-b-c", "-")) // [a b c]
	// 小写
	fmt.Println(strings.ToLower(a)) // hello
	// 大写
	fmt.Println(strings.ToUpper(a)) // HELLO
	// 长度
	fmt.Println(len(a)) // 5
	b := "你好"
	fmt.Println(len(b)) // 6
}
