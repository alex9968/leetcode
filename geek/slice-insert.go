package main

import (
	"errors"
	"fmt"
)

func main() {
	values  := []interface{}{0,1,2,3,5}

	newValues,err  := Insert(values, 4, 0)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	fmt.Printf("%#v\n", newValues )

	newValues2 := Delete(newValues, 0)
	fmt.Printf("%#v", newValues2 )


}

//Insert element to slice
func Insert(values []interface{}, val interface{}, index int)([]interface{}, error) {
	res := []interface{}{}

	if index <  0 || index > len(values) {
		return nil, errors.New("index error")
	}

	for i := 0; i < index; i++ {
		res = append(res, values[i])
	}

	res = append(res, val)
	for i := index; i < len(values); i++ {
		res = append(res, values[i])
	}
	return res ,  nil
}

//Delete some one
func Delete(values []interface{}, index int ) interface{}{
	if index <0 || index > len(values) {
		return nil
	}

	res := []interface{}{}
	for i := 0; i < len(values); i++ {
		if index == i {
			continue
		}
		 res = append(res, values[i])
	}
	return res
}

