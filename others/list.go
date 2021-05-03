package main

import (
	"fmt"
)

type Boy struct {
	No int 
	Next *Boy 
}

func AddBoy(num int) *Boy {
	first := &Boy{}

	curBoy := &Boy{}

	if num < 1  {
		fmt.Println("num err")
		return first
	}

	for i := 1; i<= num; i++ {
		boy := &Boy{
			No: 1
		}

		if i == 1 {
			first = boy 
			curBoy = boy
			curBoy.Next = first
		}else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}
	return first
}

func ShowBoy(first *Boy) {
	if first.Next == nil {
		fmt.Println("boy null")
		return
	}

	curBoy := first

	for {
		fmt.Println("boy number: %d ->", curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
}

func  PlayGame(first &Boy, startNo int, countNUm int) {
	
}