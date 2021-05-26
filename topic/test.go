package main

import (
	"fmt"
)

//channel关闭之后，仍然可以从channel中读取剩余的数据，直到数据全部读取完成
func main() {

	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)
	//type 1
	for value := range ch {
		fmt.Println("value:", value)
	}
	//type 2
	for {
		select {
		case a := <-ch:
			fmt.Println("a", a)
		}
	}
}

func f() {

}

func f() int {
	// fmt.Print(a)
	return 0
}
