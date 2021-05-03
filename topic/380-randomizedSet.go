package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type RandomizedSet struct {
	m    map[string]int
	nums []int
}

func main() {
	obj := Constructor()
	ok1 := obj.Insert(1)
	ok2 := obj.Insert(2)
	ok3 := obj.Insert(3)

	fmt.Printf("insert: %v,  %v, %v \n", ok1, ok2, ok3)
	obj.Show()

	ok4 := obj.Remove(2)
	fmt.Printf("remove:  %v \n", ok4)
	obj.Show()
	num := obj.GetRandom()

	fmt.Printf("get random:  %v \n", num)
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().Unix())
	return RandomizedSet{
		m:    make(map[string]int),
		nums: []int{},
	}
}

func (this *RandomizedSet) Show() {
	fmt.Println(this.nums)
}

func (this *RandomizedSet) Insert(val int) bool {
	key := strconv.Itoa(val) //to string
	if _, ok := this.m[key]; ok {
		return false
	} //存在直接返回false

	this.nums = append(this.nums, val)
	index := len(this.nums) - 1
	this.m[key] = index //存入index
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	key := strconv.Itoa(val)

	if _, ok := this.m[key]; !ok {
		return false
	}

	index := this.m[key] //获得index
	delete(this.m, key)  //删除

	num := this.nums[len(this.nums)-1]       //get last
	this.nums = this.nums[:len(this.nums)-1] //delete last

	if num != val {
		this.nums[index] = num            //把最后一个误删的放入空位
		this.m[strconv.Itoa(num)] = index //update map
	}
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.nums))
	return this.nums[index]
}
