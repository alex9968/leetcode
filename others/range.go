package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

//用户
func main() {

	a := []*Person{
		{"tom", 19},
		{"tom2", 29},
		{"tom3", 39},
	}

	for i, v := range a {
		fmt.Printf("%d: %p\n", i, v)
	}

	f()
}

func f() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, val := range slice {
		m[key] = &val //val的地址一直都没变，是同一个，他最后的值为3
	}
	for k, v := range m {
		fmt.Println(k, "->", *v) //都是3
	}
}
