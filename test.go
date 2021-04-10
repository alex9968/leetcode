package main

import "fmt"

func main() {
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	fmt.Printf("%v \n", pairs[')'] > 0)
	fmt.Printf("%v \n", pairs[']'] > 0)
	fmt.Printf("%v \n", pairs['}'] > 0)
	fmt.Printf("%v \n", pairs['0'])
}
