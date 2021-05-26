package main

import (
	"fmt"
)

//用户
func main() {

	a := []int{}
	fmt.Printf("11 %p\n", &a)
	a = f(a)
	a = f(a)
	a = f(a)
	a = f(a)
	a = f(a)
	a = f(a)
	fmt.Print(a)
}

func f(s []int) []int {

	// fmt.Printf("ff %p\n", &s) //地址会变因为slice底层是结构体，结构体是值类型

	s = append(s, 2)
	fmt.Printf("Addr:%p, len(%d), cap(%d)\n,", &s, len(s), cap(s))
	return s
}
