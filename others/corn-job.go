package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func Test1()  {
	time.Sleep(10*time.Second)
	fmt.Println("Test1...")
}
func Test2()  {
	fmt.Println("Test2...")
}
func main() {
	fmt.Println("starting go cron...")

	i := 0
	cron := cron.New()
	spec1 := "*/1 * * * * ?"
	//直接闭包函数
	cron.AddFunc(spec1, func() {
		i++
		fmt.Println("cron is running...", i)
	})

	spec2 := "*/1 * * * * ?"
	cron.AddFunc(spec1, Test1) //运行函数Test1
	cron.AddFunc(spec2, Test2) //运行函数Test2

	//启动计划任务
	cron.Start()

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer cron.Stop()

	select {}
}
