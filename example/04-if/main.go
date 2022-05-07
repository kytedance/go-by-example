package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// 判断闰年
	// 如果年份能被4整除，但不能被100整除，或者能被400整除，则是闰年
	if year := 2022; year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Println("\n", year, "is a leap year")
	} else {
		fmt.Println("\n", year, "is not a leap year")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
