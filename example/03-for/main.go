package main

import "fmt"

func main() {
	// 输出从 1 到 100
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
	fmt.Println("------------")
	// 输出 1 - 100 的奇数
	for i := 1; i <= 100; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
	fmt.Print("------------\n")

	// i := 1
	// for {
	// 	fmt.Println("loop")
	// 	break
	// }
	// for j := 7; j < 9; j++ {
	// 	fmt.Println(j)
	// }

	// for n := 0; n < 5; n++ {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }
}
