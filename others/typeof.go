package main

import (
	"fmt"
	"reflect"
)

func main() {

	t := reflect.TypeOf(f)

	// f()

	// fmt.Print(t.Name())

	aa, _ := t.MethodByName("f")
	fmt.Println("aa:", aa.Name)

}

func f() {
	fmt.Println("ff")
}
