// var isValid = function(s) {

//     const maping = new Map();
//     maping.set('(', ')');
//     maping.set('[', ']');
//     maping.set('{', '}');

//     let stack = [];

//     for (let i =0; i< s.length; i++){
//         if(maping.has(s[i])) {
//             // 如果是左半边， 那么放到stack中去，左半边可以无限多
//             stack.push(maping.get(s[i]))
//         }else {
//             if(stack.pop() !== s[i])  //如果是右半边则必须和stack中的左半边配对
//             return false
//         }
//     }

//     if(stack.length) return false
//     return true
// };

package main

import "fmt"

func main() {

	fmt.Printf("%v \n", isValid("()[]{}"))  //true
	fmt.Printf("%v \n", isValid("())[]{}")) //false

}

func isValid(s string) bool {

	n := len(s)
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}

	fmt.Printf("%v \n", pairs[')'] > 0) // true
	fmt.Printf("%v \n", pairs[']'] > 0) // true
	fmt.Printf("%v \n", pairs['}'] > 0) // true
	fmt.Printf("%v \n", pairs['0'] > 0) // false, pairs['.'] == 0

	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 { //是 ）] } ,说明该配对了
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] { //没东西配、 没配上 , 直接false
				return false
			}
			stack = stack[:len(stack)-1] //到这里说明配对了， 就可以pop一个出来了
		} else {
			//不需要配对，直接塞进stack
			stack = append(stack, s[i])
		}

	}
	return len(stack) == 0
}
