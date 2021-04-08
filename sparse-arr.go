package main

import (
	"fmt"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//create arr
	var chessMap [6][6]int
	chessMap[1][2] = 1 //black node
	chessMap[2][3] = 2 //white node

	fmt.Println("原始数组:")
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//传唤为稀疏数组
	var sparseArr []ValNode

	valNode := ValNode{
		row: 6,
		col: 6,
		val: 0,
	}

	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	fmt.Println("当前稀疏数组：")
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d \n", i, valNode.row, valNode.col, valNode.val)
	}

	//恢复原来的数组
	var chessMap2 [6][6]int
	for i, valNode := range sparseArr {
		if i != 0 {
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}

	//校验chessMap2是否恢复
	fmt.Println("恢复后的数据：")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

}
