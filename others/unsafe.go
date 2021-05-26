package main

import (
	"fmt"
	"unsafe"
)

//用户
func main() {
	a := [4]int{0, 1, 2, 3}
	fmt.Println("Arr Sizeof:", unsafe.Sizeof(a))   //整体大小
	fmt.Println("Arr Alignof:", unsafe.Alignof(a)) //对齐粒度
	p1 := unsafe.Pointer(&a[1])
	//开始内存偏移
	p3 := (*int)(unsafe.Pointer(uintptr(p1) + 2*unsafe.Sizeof(a[0]))) //让指针移动到a[3]
	*p3 = 100
	fmt.Println("a =", a) // a = [0 1 2 100]

	fmt.Println()

	type Person struct {
		name   string
		name2  string
		age    int
		gender bool
	}

	who := Person{"John", "22", 30, true}
	fmt.Println("Person Sizeof:", unsafe.Sizeof(who))   //整体大小
	fmt.Println("Person Alignof:", unsafe.Alignof(who)) //对齐粒度
	pp := unsafe.Pointer(&who)
	//开始内存偏移, Offsetof偏移到结构体某个字段处
	pname := (*string)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.name)))
	pname2 := (*string)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.name) + unsafe.Sizeof("")))
	page := (*int)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.age)))
	*pname = "Alice"
	*pname2 = "Alice22"
	*page = 28
	fmt.Println(who) // {Alice 28 true}

}
