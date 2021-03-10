package data

import (
	"fmt"
	"testing"
)

/*
	字符串：不可变的字节序列
	默认UTF-8编码，存储Unicode字符，字面量允许使用十六进制，八进制和UTF编码格式
*/
func TestString(t *testing.T) {
	s := "震宇\x61\142\u0041"

	fmt.Printf("%s\n", s)
	fmt.Printf("% x, len: %d \n", s, len(s))
}
