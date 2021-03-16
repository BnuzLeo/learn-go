package data

import (
	"fmt"
	"reflect"
	"testing"
)

func printSlice(s interface{}, name string) {
	arr := reflect.ValueOf(s)
	if arr.Kind() != reflect.Slice {
		fmt.Printf("%v it si not a slice", name)
		return
	}
	fmt.Printf("%v=%v, len(%v)=%v, cap(%v)=%v\n", name, s, name, arr.Len(), name, arr.Cap())
}

/*
	Slice的声明
*/

func TestCreateSlice(t *testing.T) {
	//直接声明
	var s []int
	printSlice(s, "s")
	//声明 + 赋值
	s1 := []int{1, 2, 3}
	printSlice(s1, "s1")
	//声明 + 开空间
	s2 := make([]int, 10)
	printSlice(s2, "s2")
}

/*
	Slice的基本用法
*/
func updateSlice(s []int) {
	if len(s) > 3 {
		s[2] = 100
	}
}

func TestSlice(t *testing.T) {
	/*
		arr[2:5] : [2 3 4]
		arr[2:] : [2 3 4 5 6 7 8]
		arr[:5] : [0 1 2 3 4]
		arr[:] : [0 1 2 3 4 5 6 7 8]
	*/
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("arr[2:5] : %v\n", arr[2:5])
	fmt.Printf("arr[2:] : %v\n", arr[2:])
	fmt.Printf("arr[:5] : %v\n", arr[:5])
	fmt.Printf("arr[:] : %v\n", arr[:])
}

/*
	len和cap：

	slice可以往后扩展，不可以向前扩展
	s[i]不可以超越len，向后扩展不可以超越cap
*/

func TestLenAndCap(t *testing.T) {
	s := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s=%v, len(s)=%v, cap(s)=%v\n", s, len(s), cap(s))
	s1 := s[2:6]
	fmt.Printf("s1=%v, len(s1)=%v, cap(s1)=%v\n", s1, len(s1), cap(s1))
	s2 := s1[0:5]
	fmt.Printf("s2=%v, len(s2)=%v, cap(s2)=%v\n", s2, len(s2), cap(s2))
}

/*
	添加元素的时候，如果超越了cap，系统会重新分配更大的底层数组，然后把他拷贝过去

	由于是值传递，所以一定要接收append的返回值
*/
func TestAppend(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	s1 := append(s, 6)
	s2 := append(s1, 7)
	s3 := append(s2, 8)
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
}

/*
	操作slice
*/
func TestOperation(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	printSlice(s, "s")
	// 删除某个元素
	s = append(s[:2], s[3:]...)
	printSlice(s, "s")
	// 顶部弹出
	tail := s[len(s)-1]
	s = s[:len(s)-1]
	printSlice(s, "s")
	fmt.Println(tail)
	//底部弹出
	front := s[0]
	s = s[1:]
	printSlice(s, "s")
	fmt.Println(front)
}
