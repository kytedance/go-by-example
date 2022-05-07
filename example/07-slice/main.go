package main

import "fmt"

func main() {
	// 可变长度的切片
	// s := make([]string, 3)
	myS := make([]string, 5)
	myS[0] = "L"
	myS[1] = "I"
	myS[2] = "J"
	myS[3] = "I"
	myS[4] = "E"
	fmt.Println("s:", myS)

	myS = append(myS, "I", "S", "H", "A", "N", "D", "S", "O", "M", "E")
	fmt.Println("s:", myS)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"
	// fmt.Println("get:", s[2])   // c
	// fmt.Println("len:", len(s)) // 3

	fmt.Println("len:", len(myS)) // 15
	// 左闭右开区间
	fmt.Println(myS[:5])
	fmt.Println(myS[5:7])
	fmt.Println(myS[7:])
	// s = append(s, "d")
	// s = append(s, "e", "f")
	// fmt.Println(s) // [a b c d e f]

	// c := make([]string, len(s))
	// copy(c, s)
	// fmt.Println(c) // [a b c d e f]

	// fmt.Println(s[2:5]) // [c d e]
	// fmt.Println(s[:5])  // [a b c d e]
	// fmt.Println(s[2:])  // [c d e f]

	// good := []string{"g", "o", "o", "d"}
	// fmt.Println(good) // [g o o d]
}
