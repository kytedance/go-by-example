package main

import "fmt"

func main() {
	// 快捷遍历操作

	myNums := []string{"I", "A", "M", "H", "A", "P", "P", "Y"}
	for index, value := range myNums {
		fmt.Println(index, value)
	}
	myMaps := map[string]string{"name": "happy", "age": "18"}
	for key, value := range myMaps {
		fmt.Println(key, value)
	}
	for _, value := range myMaps {
		fmt.Println(value)
	}
	for key, _ := range myMaps {
		fmt.Println(key)
	}

	// 	nums := []int{2, 3, 4}
	// 	sum := 0
	// 	for i, num := range nums {
	// 		sum += num
	// 		if num == 2 {
	// 			fmt.Println("index:", i, "num:", num) // index: 0 num: 2
	// 		}
	// 	}
	// 	fmt.Println(sum) // 9

	// 	m := map[string]string{"a": "A", "b": "B"}
	// 	for k, v := range m {
	// 		fmt.Println(k, v) // b 8; a A
	// 	}
	// 	for k := range m {
	// 		fmt.Println("key", k) // key a; key b
	// 	}
}
