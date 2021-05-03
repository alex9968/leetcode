package main

func main() {

}

func Test(s string) int {

	res, index := 0, -1

	m := map[byte]int{}
	n := len(s)

	for i:= 0; i<n; i++ {
		if i!=0{
			delete(m, [i-1])
		}

		for index+1 <n && m[s[index+1]] ==0 {
			m[s[index+1]]++
			index++
		}
		res = max(res, index-i+1)
	} 
	return res

}
