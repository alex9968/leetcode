package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
}

type Tow struct {
	Person1 Person `json:"person1"`
	Person2 Person `json:"person2"`
}

//channel关闭之后，仍然可以从channel中读取剩余的数据，直到数据全部读取完成
func main() {
	t := Tow{
		Person1 : Person {
		    Name: "sss",
		},
		Person2 : Person {
			Name: "sss22",
		},
	}

	j, err := json.Marshal(t)
	if err != nil {
		fmt.Println("marshal err")
	}

	t2 := Tow{}
	err = json.Unmarshal(j, &t2)
	if err != nil {
		fmt.Println("Unmarshal err")
	}

	fmt.Println(" ", t2)
}



