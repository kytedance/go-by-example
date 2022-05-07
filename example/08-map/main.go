package main

import "fmt"

func main() {
	// golang map 完全无序
	// 构造一个map
	myMap := make(map[string]int)

	myMap["code"] = 2001
	myMap["status"] = 101
	fmt.Println("myMap:", myMap)

	fmt.Println("len:", len(myMap))

	fmt.Println("get_code:", myMap["code"])

	fmt.Println("get_status:", myMap["status"])

	fmt.Println("get_not_exist:", myMap["not_exist"])

	// 删除一个键值对
	delete(myMap, "status")

	fmt.Println("len:", len(myMap))

	// 判断是否存在某个键
	_, ok := myMap["status"]
	fmt.Println("ok:", ok)
	// m := make(map[string]int)
	// m["one"] = 1
	// m["two"] = 2
	// fmt.Println(m)           // map[one:1 two:2]
	// fmt.Println(len(m))      // 2
	// fmt.Println(m["one"])    // 1
	// fmt.Println(m["unknow"]) // 0

	// r, ok := m["unknow"]
	// fmt.Println(r, ok) // 0 false

	// delete(m, "one")

	// m2 := map[string]int{"one": 1, "two": 2}
	// var m3 = map[string]int{"one": 1, "two": 2}
	// fmt.Println(m2, m3)
}
