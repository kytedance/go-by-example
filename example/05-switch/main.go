package main

import (
	"fmt"
	"time"
)

func main() {
	// golang switch 只会匹配一次
	weekDay := time.Now().Weekday()
	switch weekDay {
	case time.Saturday:
		fmt.Println("Today is Saturday")
	case time.Sunday:
		fmt.Println("Today is Sunday")
	case time.Monday:
		fmt.Println("Today is Monday")
	case time.Tuesday:
		fmt.Println("Today is Tuesday")
	case time.Wednesday:
		fmt.Println("Today is Wednesday")
	case time.Thursday:
		fmt.Println("Today is Thursday")
	case time.Friday:
		fmt.Println("Today is Friday")
	default:
		fmt.Println("Today is unknown")
	}

	// a := 2
	// switch a {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// case 4, 5:
	// 	fmt.Println("four or five")
	// default:
	// 	fmt.Println("other")
	// }

	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("It's before noon")
	// default:
	// 	fmt.Println("It's after noon")
	// }
}
