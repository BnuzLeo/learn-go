package data

import (
	"fmt"
	"testing"
)

/*
	Map的基本操作
*/
func TestMapBase(t *testing.T) {
	// 初始化声明
	m1 := map[string]string{
		"name":    "leon",
		"address": "gd",
	}
	fmt.Println(m1)
	// empty map声明
	m2 := make(map[string]string)
	fmt.Println(m2)
	// nil map声明
	var m3 map[string]string
	fmt.Println(m3)

	// 获取单个元素
	name, ok := m1["name"]
	fmt.Println(name, ok)

	// 删除单个key
	delete(m1, "name")
	fmt.Println(m1)
}

/*
	算法：找出一个字符串中，最大不重复的字符串长度
	如："bbbb":1, "abcdef":6, "abcab":3
*/
func TestMaxLen(t *testing.T) {
	fmt.Println(maxLen("bbbb"))
	fmt.Println(maxLen("abcdef"))
	fmt.Println(maxLen("abcab"))
}

func maxLen(str string) int {
	occur := make(map[byte]int)
	len := 0
	start := 0

	bytes := []byte(str)
	for i, b := range bytes {
		if idx, ok := occur[b]; ok && idx >= start {
			start = idx + 1
		}
		if i-start+1 > len {
			len = i - start + 1
		}
		occur[b] = i
	}
	return len
}
