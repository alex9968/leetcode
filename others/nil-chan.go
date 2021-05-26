package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 不断向channel c中发送[0,10)的随机数
// 结论： nil chan 发送接收都会阻塞，close会panic
func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func get(c chan int) {
	sum := 0
	// 1秒后，将向t.C通道发送时间点，使其可读
	t := time.NewTimer(1 * time.Second)
	// 一秒内，将一直选择第一个case
	// 一秒后，t.C可读，将选择第二个case
	// c变成nil channel后，两个case分支都将一直阻塞
	for {
		select {
		case num := <-c:
			sum += num
		case <-t.C: //一秒后接收到信号
			c = nil
			fmt.Println("sum:", sum)
		}
	}
}

func main() {
	c := make(chan int)
	go get(c)
	go send(c)
	// 给3秒时间让前两个goroutine有足够时间运行
	time.Sleep(3 * time.Second)
}
