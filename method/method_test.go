package method

import (
	"fmt"
	"testing"
)

/*
	Go的函数是一等公民：
	Go的函数是第一对象，可以用作函数的入参和返回值，也可以存入对象实体
*/

func format(str string) (format string) {
	result := str + " format"
	fmt.Println(result)
	return result
}

func TestFirstObject(t *testing.T) {
	Format(format, "first object")
}
