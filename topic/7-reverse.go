package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("res", reverse(123))
	fmt.Println("res", reverse(300))

}

//整数反转

// 时间复杂度：O(\log |x|)O(log∣x∣)。翻转的次数即 xx 十进制的位数。
// 空间复杂度：O(1)
func reverse(x int) (res int) {
	for x != 0 {
		//进入判断即认为，至少还有一个位，所以res*10<MinInt32
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		res = res*10 + digit
	}
}

func f(x int) (res int) {
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}

		digit := x % 10
		x /= 10
		res = res*10 + digit
	}
}
