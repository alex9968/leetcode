package main

import (
	"fmt"
)

func main() {
	var a interface{} = 10
	var b int = 10

	fmt.Print(a == b) //true ,类型和值都相等才行

	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}

}

func ParseBool(val interface{}) (value bool, err error) {
	if val != nil {
		switch v := val.(type) {
		case bool:
			return v, nil
		case string:
			switch v {
			case "1", "t", "T", "true", "TRUE", "True", "YES", "yes", "Yes", "Y", "y", "ON", "on", "On":
				return true, nil
			case "0", "f", "F", "false", "FALSE", "False", "NO", "no", "No", "N", "n", "OFF", "off", "Off":
				return false, nil
			}
		case int8, int32, int64:
			strV := fmt.Sprintf("%s", v)
			if strV == "1" {
				return true, nil
			} else if strV == "0" {
				return false, nil
			}
		case float64:
			if v == 1 {
				return true, nil
			} else if v == 0 {
				return false, nil
			}
		}
		return false, fmt.Errorf("parsing %q: invalid syntax", val)
	}
	return false, nil
}

// type Addable interface {
// 	type int, int8, int16, int32, int64,
// 		uint, uint8, uint16, uint32, uint64, uintptr,
// 		float32, float64, complex64, complex128,
// 		string
// }
// func add[T Addable](a, b T) T {
// 	return a + b
// }
// func main() {
// 	fmt.Println(add(1,2))
// 	// FIXME
// 	//fmt.Println(add("foo","bar"))
// }