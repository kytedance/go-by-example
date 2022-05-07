package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
	// 每次生成的随机数都一样，需要生成随机数种子
}
