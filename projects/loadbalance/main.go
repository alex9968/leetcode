package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/alex9968/leetcode/loadbalance/balance"
)

func main() {

	// 创建instance
	var insts []*balance.Instance
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		insts = append(insts, balance.NewInstance(host, 8080))
	}

	var balanceName = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	for {
		//每秒选出一个
		inst, err := balance.DoBalance(balanceName, insts)
		if err != nil {
			fmt.Println("do balance err:", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}

}
