package main

import (
	"errors"
	"fmt"
)

type user struct {
	name     string
	password string
}

// bool      -> false
// numbers   -> 0
// string    -> ""
// pointers -> nil
// slices -> nil
// maps -> nil
// channels -> nil
// functions -> nil
// interfaces -> nil

// 直接在函数返回值中返回错误
func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			// 没有错误返回 nil
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func main() {
	u, err := findUser([]user{{"wang", "1024"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name) // wang

	if u, err := findUser([]user{{"wang", "1024"}}, "li"); err != nil {
		fmt.Println(err) // not found
		return
	} else {
		fmt.Println(u.name)
	}
}
