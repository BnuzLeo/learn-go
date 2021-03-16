package data

import (
	"fmt"
	"testing"
)

/*
	数组的基本用法
*/
func TestArrayBaseUse(t *testing.T) {
	var a [5]int
	a2 := [3]int{1, 3, 5}
	a3 := [...]int{1, 2, 3, 4, 5, 6} // 让编译器来帮你算长度
	fmt.Println(a, a2, a3)

	grid := [4][5]int{}
	fmt.Println(grid)
}

/*
	数组是值类型
	这里的[5]int是数组，如果是[]int是切片
*/
func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func changArr(arr [5]int) {
	arr[0] = 100
}

func TestArrayTransfer(t *testing.T) {
	arr := [5]int{}
	changArr(arr) // 值并不会边
	printArray(arr)
}
