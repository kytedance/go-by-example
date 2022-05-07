package main

import (
	"fmt"
	"math"
)

func main() {

	var a = "initial"
	var my_a = "my_initial"

	var b, c int = 1, 2
	var myB, myC int = 3, 4

	var d = true
	var myD = false

	var e float64 = 1.1111
	var myE float64 = 2.2222

	f := float32(e)
	myF := float32(myE)

	g := a + "foo"
	myG := my_a + "foo"

	fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
	fmt.Println(g)                // initialapple

	fmt.Println(my_a, myB, myC, myD, myE, myF) // my_initial 3 4 false 0 0
	fmt.Println(myG)                           // my_initialapple

	const s string = "constant"
	const myS string = "my_constant"

	const h = 500000000
	const myH = 500000001

	const i = 3e20 / h
	const myI = 3e20 / myH

	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))

	fmt.Println(myS, myH, myI, math.Sin(myH), math.Sin(myI))
}
