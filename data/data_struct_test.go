package data

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
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

/*
	string 和 []byte 和 []rune的转换
	所以在转换的时候，无论怎么转，都会分配内存，重新拷贝
*/

func pp(format string, prt interface{}) {
	pointer := reflect.ValueOf(prt).Pointer()
	h := (*uintptr)(unsafe.Pointer(pointer))
	fmt.Printf(format, h)
}

func TestStringConvertByteAndRuneArray(t *testing.T) {
	s := "震宇"
	pp("string's address: %x\n", &s)

	bytes := []byte(s)
	bytesToStr := string(bytes)
	pp("string to bytes, bytes' address : %x\n", &bytes)
	pp("bytes to string, string' address : %x\n", &bytesToStr)

	runes := []rune(s)
	runesToStr := string(runes)
	pp("string to runes, runes' address : %x\n", &runes)
	pp("runes to string, string' address : %x\n", &runesToStr)
}

/*
	如何做到转换不分配内存
*/

func TestDontCreateMemory(t *testing.T) {
	s := "震宇"

	bs := []byte(s)
	bs = append(bs, "I lover you"...)

	result := (*string)(unsafe.Pointer(&bs))
	pp("string to []byte, []byte address: %x\n", &bs)
	pp("[]byte to string, string address: %x\n", result)
}

/*
	byte和rune的区别
*/
func TestByteAndRune(t *testing.T) {
	s := "震宇"

	/*
		0 [é]
		1 []
		2 []
		3 [å]
		4 [®]
		5 []
	*/
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d [%c]\n", i, s[i])
	}

	/*
		0 [震]
		3 [宇]
	*/
	for i, c := range s {
		fmt.Printf("%d [%c]\n", i, c)
	}
}

/*
	拼接字符串的消耗
	解决方法：
		方法1: strings.Join
		方法2：bytes.Buffer
*/
func add() string {
	var s string
	for i := 0; i < 1000; i++ {
		s += "a"
	}
	return s
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add() //Benchmark-8   	    9949	    116630 ns/op
	}
}

func join() string {
	s := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = "a"
	}
	return strings.Join(s, "")
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		join() // BenchmarkJoin-8   	  139749	      8296 ns/op
	}
}

func buff() string {
	var b bytes.Buffer
	b.Grow(1000)
	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}
	return b.String()
}

func BenchmarkBuff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buff() //BenchmarkBuff-8   	  228507	      5124 ns/op
	}
}
